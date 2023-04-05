package service

import (
	"project-tiga/helper"
	"project-tiga/model"
	"project-tiga/repository"
)

type OrderService struct {
	OrderRepository repository.IOrderRepository
}

func NewOrderService(orderRepository repository.IOrderRepository) *OrderService {
	return &OrderService{
		OrderRepository: orderRepository,
	}
}

func (os *OrderService) Create(request model.OrderCreateRequest, userID string) (model.OrderCreateResponse, error) {
	id := helper.GenerateID()

	order := model.Order{
		ID:     id,
		UserID: userID,
		Price:  request.Price,
	}

	err := os.OrderRepository.Add(order)
	return model.OrderCreateResponse{
		ID:     id,
		UserID: userID,
		Price:  request.Price,
	}, err
}

func (os *OrderService) GetListOrders(userID string) ([]model.OrderListResponse, error) {
	orders, err := os.OrderRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return os.toOrderListResponseFrom(orders), nil
}

func (os *OrderService) toOrderListResponseFrom(orders []model.Order) []model.OrderListResponse {
	orderListResponses := make([]model.OrderListResponse, 0)

	for _, order := range orders {
		orderListResponses = append(orderListResponses, model.OrderListResponse{
			ID:    order.ID,
			Price: order.Price,
		})
	}

	return orderListResponses
}
