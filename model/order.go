package model

type Order struct {
	ID     string `gorm:"primaryKey;type:varchar(255)"`
	Price  int    `gorm:"not null"`
	UserID string
}

type OrderCreateRequest struct {
	Price int `json:"price" valid:"required~price is blank"`
}

type OrderCreateResponse struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Price  int    `json:"price"`
}

type OrderListResponse struct {
	ID    string `json:"id"`
	Price int    `json:"price"`
}
