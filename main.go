package main

import (
	"project-tiga/controller"
	"project-tiga/middleware"
	"project-tiga/model"
	"project-tiga/repository"
	"project-tiga/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error

	db, err = gorm.Open(postgres.Open("host=containers-us-west-178.railway.app port=6798 user=postgres password=VAiHbJpUTI5uMIK3WwNG dbname=railway sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(model.User{}, model.Order{})
}

func main() {
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(*orderService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)
	userController := controller.NewUserController(*userService)

	g := gin.Default()

	g.POST("/user/register", userController.Register)
	g.POST("/user/login", userController.Login)
	g.POST("/user/refresh", middleware.AuthRefreshMiddleware, userController.Refresh)

	orderGroup := g.Group("/order", middleware.AuthMiddleware)
	orderGroup.POST("/", orderController.CreateOrder)
	orderGroup.GET("/", orderController.GetListOrders)
	orderGroup.GET("/:id", orderController.GetOrder)

	err := g.Run(":8084")
	if err != nil {
		panic(err)
	}
}
