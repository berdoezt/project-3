package main

import (
	"fmt"
	"os"
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
	db   *gorm.DB
	PORT = os.Getenv("PORT")

	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
)

func init() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	err := g.Run(":" + PORT)
	if err != nil {
		panic(err)
	}
}
