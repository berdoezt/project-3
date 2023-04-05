package service

import (
	"project-tiga/model"
	"project-tiga/repository"
	"project-tiga/repository/mocks"
	"reflect"
	"testing"
)

func TestOrderService_GetListOrders(t *testing.T) {

	orderRepository := mocks.NewIOrderRepository(t)

	type fields struct {
		OrderRepository repository.IOrderRepository
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     []model.OrderListResponse
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "Case #1",
			fields: fields{
				OrderRepository: orderRepository,
			},
			args: args{
				userID: "1",
			},
			want: []model.OrderListResponse{
				{
					ID:    "1",
					Price: 100,
				},
				{
					ID:    "2",
					Price: 200,
				},
			},
			mockFunc: func() {
				orderRepository.On("GetByUserID", "1").Return([]model.Order{
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
				}, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			os := &OrderService{
				OrderRepository: tt.fields.OrderRepository,
			}

			tt.mockFunc()

			got, err := os.GetListOrders(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.GetListOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.GetListOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
