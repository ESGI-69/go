/*
Package broadcast provides pubsub of messages over channels.

A provider has a Broadcaster into which it Submits messages and into
which subscribers Register to pick up those messages.
*/
package broadcaster

import (
	"go/src/payment"
)

type broadcaster struct {
	input chan payment.Payment
	reg   chan chan<- interface{}
	unreg chan chan<- interface{}

	outputs map[chan<- interface{}]bool
}

// The Broadcaster interface describes the main entry points to
// broadcasters.
type Broadcaster interface {
	// Register a new channel to receive broadcasts
	Register(chan<- interface{})
	// Unregister a channel so that it no longer receives broadcasts.
	Unregister(chan<- interface{})
	// Shut this broadcaster down.
	Close() error
	// Try Submit a new object to all subscribers return false if input chan is fill
	Submit(payment.Payment) bool
}

func (bForBroctaster *broadcaster) broadcast(payment payment.Payment) {
	for ch := range bForBroctaster.outputs {
		ch <- payment
	}
}

func (bForBroctaster *broadcaster) run() {
	for {
		select {
		case message := <-bForBroctaster.input:
			bForBroctaster.broadcast(message)
		case ch, ok := <-bForBroctaster.reg:
			if ok {
				bForBroctaster.outputs[ch] = true
			} else {
				return
			}
		case ch := <-bForBroctaster.unreg:
			delete(bForBroctaster.outputs, ch)
		}
	}
}

// NewBroadcaster creates a new broadcaster with the given input
// channel buffer length.
func NewBroadcaster(buflen int) Broadcaster {
	bForBroctaster := &broadcaster{
		input:   make(chan payment.Payment, buflen),
		reg:     make(chan chan<- interface{}),
		unreg:   make(chan chan<- interface{}),
		outputs: make(map[chan<- interface{}]bool),
	}

	go bForBroctaster.run()

	return bForBroctaster
}

func (bForBroctaster *broadcaster) Register(newch chan<- interface{}) {
	bForBroctaster.reg <- newch
}

func (bForBroctaster *broadcaster) Unregister(newch chan<- interface{}) {
	bForBroctaster.unreg <- newch
}

func (bForBroctaster *broadcaster) Close() error {
	close(bForBroctaster.reg)
	close(bForBroctaster.unreg)
	return nil
}

// TrySubmit attempts to submit an item to be broadcast, returning
// true iff it the item was broadcast, else false.
func (bForBroctaster *broadcaster) Submit(payment payment.Payment) bool {
	if bForBroctaster == nil {
		return false
	}
	select {
	case bForBroctaster.input <- payment:
		return true
	default:
		return false
	}
}
