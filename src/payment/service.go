package payment

type Service interface {
	Store(input InputPayment) (Payment, error)
	FetchAll() ([]Payment, error)
	FetchById(id int) (Payment, error)
	Update(id int, input InputPayment) (Payment, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input InputPayment) (Payment, error) {
	var payment Payment
	payment.ProductId = input.ProductId
	payment.PricePaid = input.PricePaid
	payment, err := s.repository.Store(payment)
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (s *service) FetchAll() ([]Payment, error) {
	payments, err := s.repository.FetchAll()
	if err != nil {
		return payments, err
	}
	return payments, nil
}

func (s *service) FetchById(id int) (Payment, error) {
	payment, err := s.repository.FetchById(id)
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (s *service) Update(id int, input InputPayment) (Payment, error) {
	updatePayment, err := s.repository.Update(id, input)
	if err != nil {
		return updatePayment, err
	}
	return updatePayment, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
