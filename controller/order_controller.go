package controller

import (
	"net/http"
	"project-tiga/model"
	"project-tiga/service"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

// CreateOrder godoc
//
//	@Summary		create order
//	@Description	create order for a particular user
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.OrderCreateRequest	true	"request is required"
//	@Success		200		{object}	model.OrderCreateResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		BearerAuth
//	@Router			/order [post]
func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var request model.OrderCreateRequest
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

	userID, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorInvalidToken)
		return
	}

	order, err := oc.orderService.Create(request, userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (oc *OrderController) GetListOrders(ctx *gin.Context) {
	userID, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorInvalidToken)
		return
	}

	response, err := oc.orderService.GetListOrders(userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (oc *OrderController) GetOrder(ctx *gin.Context) {

}
