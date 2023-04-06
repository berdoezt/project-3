package controller

import (
	"net/http"
	"project-tiga/model"
	"project-tiga/service"

	"github.com/asaskevich/govalidator"
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

// Login godoc
//
//	@Summary		login user
//	@Description	login user using email and password
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.UserLoginRequest	true	"request is required"
//	@Success		200		{object}	model.UserLoginResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Router			/user/login [post]
func (uc *UserController) Login(ctx *gin.Context) {
	var request model.UserLoginRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	valid, err := govalidator.ValidateStruct(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: "tidak valid",
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

// Register godoc
//
//	@Summary		register a new user
//	@Description	register a new user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.UserRegisterRequest	true	"request is required"
//	@Success		200		{object}	model.UserRegisterResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Router			/user/register [post]
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
