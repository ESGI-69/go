package payment

type Service interface {
	Create(input InputPayment) (Payment, error)
	GetAll() ([]Payment, error)
	GetById(id int) (Payment, error)
	Update(id int, input InputPayment) (Payment, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input InputPayment) (Payment, error) {
	var payment Payment
	payment.ProductId = input.ProductId
	payment.PricePaid = input.PricePaid
	payment, err := s.repository.Create(payment)
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (s *service) GetAll() ([]Payment, error) {
	payments, err := s.repository.GetAll()
	if err != nil {
		return payments, err
	}
	return payments, nil
}

func (s *service) GetById(id int) (Payment, error) {
	payment, err := s.repository.GetById(id)
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
