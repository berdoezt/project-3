package controller

import (
	"net/http"
	"project-tiga/model"
	"project-tiga/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) Refresh(ctx *gin.Context) {
	userID, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorInvalidToken)
		return
	}

	response, err := uc.UserService.Refresh(userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (uc *UserController) Login(ctx *gin.Context) {
	var request model.UserLoginRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	response, err := uc.UserService.Login(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (uc *UserController) Register(ctx *gin.Context) {
	var request model.UserRegisterRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	id, err := uc.UserService.Register(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.UserRegisterResponse{
		ID: id,
	})

}
