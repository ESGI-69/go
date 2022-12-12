package payment

import "time"

type Payment struct {
	ID        int
	pricePaid float64
	createdAt time.Time
	updatedAt time.Time
}
