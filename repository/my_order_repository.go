package repository

import (
	"project-tiga/model"
)

type MyOrderRepository struct {
}

func NewMyOrderRepository() *MyOrderRepository {
	return &MyOrderRepository{}
}

func (mor *MyOrderRepository) Add(newOrder model.Order) error {

	return nil
}

func (mor *MyOrderRepository) GetByUserID(userID string) ([]model.Order, error) {

	return []model.Order{
		{
			ID:     "1",
			Price:  100,
			UserID: "1",
		},
		{
			ID:     "2",
			Price:  200,
			UserID: "1",
		},
	}, nil
}
