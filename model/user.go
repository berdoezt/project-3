package model

type User struct {
	ID       string `gorm:"primaryKey;type:varchar(255)"`
	Email    string `gorm:"not null;type:varchar(255)"`
	Password string `gorm:"not null;type:varchar(255)"`
	Orders   []Order
}

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	ID string `json:"id"`
}

type UserLoginRequest struct {
	Email    string `json:"email" valid:"required~email is blank"`
	Password string `json:"password" valid:"required~password is blank"`
}

type UserLoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
