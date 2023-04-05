package service

import (
	"project-tiga/helper"
	"project-tiga/model"
	"project-tiga/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) Register(userRegisterRequest model.UserRegisterRequest) (string, error) {
	id := helper.GenerateID()
	hashPassword, err := helper.Hash(userRegisterRequest.Password)
	if err != nil {
		return "", err
	}

	user := model.User{
		ID:       id,
		Email:    userRegisterRequest.Email,
		Password: hashPassword,
	}

	err = us.UserRepository.Add(user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (us *UserService) Login(userLoginRequest model.UserLoginRequest) (model.UserLoginResponse, error) {
	user, err := us.UserRepository.GetByEmail(userLoginRequest.Email)
	if err != nil {
		return model.UserLoginResponse{}, err
	}

	if !helper.IsHashValid(user.Password, userLoginRequest.Password) {
		return model.UserLoginResponse{}, model.ErrorInvalidEmailOrPassword
	}

	token, err := helper.GenerateAccessToken(user.ID)
	if err != nil {
		return model.UserLoginResponse{}, err
	}

	refreshToken, err := helper.GenerateRefreshToken(user.ID)
	if err != nil {
		return model.UserLoginResponse{}, err
	}

	return model.UserLoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (us *UserService) Refresh(userID string) (model.UserLoginResponse, error) {
	_, err := us.UserRepository.GetByID(userID)
	if err != nil {
		return model.UserLoginResponse{}, err
	}

	token, err := helper.GenerateAccessToken(userID)
	if err != nil {
		return model.UserLoginResponse{}, err
	}

	refreshToken, err := helper.GenerateRefreshToken(userID)
	if err != nil {
		return model.UserLoginResponse{}, err
	}

	return model.UserLoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
