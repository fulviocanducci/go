package usecase

import "github.com/fulviocanducci/go/internal/entity"

type OrderInput struct {
	Id    string
	Tax   float64
	Price float64
}

type OrderOutput struct {
	Id         string
	Tax        float64
	Price      float64
	FinalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.Id, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutput{
		Id:         order.Id,
		Tax:        order.Tax,
		Price:      order.Price,
		FinalPrice: order.FinalPrice,
	}, nil
}
