package repository

import (
	"project-tiga/model"

	"gorm.io/gorm"
)

//go:generate mockery --name IOrderRepository
type IOrderRepository interface {
	Add(newOrder model.Order) error
	GetByUserID(userID string) ([]model.Order, error)
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) Add(newOrder model.Order) error {
	tx := or.db.Create(&newOrder)
	return tx.Error
}

func (or *OrderRepository) GetByUserID(userID string) ([]model.Order, error) {
	orders := make([]model.Order, 0)
	tx := or.db.Where("user_id = ?", userID).Find(&orders)
	return orders, tx.Error
}
