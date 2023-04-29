package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"miniproject_go_wardahfdn/app/delivery"
	"miniproject_go_wardahfdn/app/infrastructure/database"
	"miniproject_go_wardahfdn/app/infrastructure/repository"
	"miniproject_go_wardahfdn/app/usecase"
)

func main() {
	db := database.InitDatabase()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)
	delivery.NewUserHandler(e, userUC)

	restaurantRepo := repository.NewRestaurantRepository(db)
	restaurantUC := usecase.NewRestaurantUseCase(restaurantRepo)
	delivery.NewRestaurantHandler(e, restaurantUC)

	menuRepo := repository.NewMenuRepository(db)
	menuUC := usecase.NewMenuUseCase(menuRepo)
	delivery.NewMenuHandler(e, menuUC)

	orderRepo := repository.NewOrderRepository(db)
	orderUC := usecase.NewOrderUseCase(orderRepo, restaurantRepo, menuRepo)
	delivery.NewOrderHandler(e, orderUC)

	paymentRepo := repository.NewPaymentRepository(db)
	paymentUC := usecase.NewPaymentUseCase(paymentRepo, orderRepo)
	delivery.NewPaymentHandler(e, paymentUC)

	log.Fatal(e.Start(":8081"))
}
