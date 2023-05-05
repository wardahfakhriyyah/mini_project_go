package delivery

import (
	"miniproject_go_wardahfdn/app/usecase"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	UserUseCase       usecase.UserUseCase
	RestaurantUseCase usecase.RestaurantUseCase
	MenuUseCase       usecase.MenuUseCase
	OrderUseCase      usecase.OrderUseCase
	PaymentUseCase    usecase.PaymentUseCase
}

func NewRoutes(h *Handler, jwtMiddleware echo.MiddlewareFunc, jsonMiddleware echo.MiddlewareFunc) *echo.Echo {
	e := echo.New()

	e.Use(jsonMiddleware)
	e.Use(jwtMiddleware)

	// User Routes
	e.GET("/users", h.GetAllUser)
	e.GET("/users/:id", h.GetUserByID)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)

	// Restaurant Routes
	e.POST("/restaurants", h.CreateRestaurant)
	e.GET("/restaurants", h.GetAllRestaurant)
	e.GET("/restaurants/:id", h.GetRestaurantByID)
	e.PUT("/restaurants/:id", h.UpdateRestaurant)
	e.DELETE("/restaurants/:id", h.DeleteRestaurant)

	// Menu Routes
	e.POST("/menus", h.CreateMenu)
	e.GET("/menus", h.GetAllMenu)
	e.GET("/menus/:id", h.GetMenuByID)
	e.PUT("/menus/:id", h.UpdateMenu)
	e.DELETE("/menus/:id", h.DeleteMenu)

	// Order Routes
	e.POST("/orders", h.CreateOrder)
	e.GET("/orders", h.GetAllOrder)
	e.GET("/orders/:id", h.GetOrderByID)
	e.PUT("/orders/:id", h.UpdateOrder)
	e.DELETE("/orders/:id", h.DeleteOrder)

	// Payment Routes
	e.POST("/payments", h.CreatePayment)
	e.GET("/payments", h.GetAllPayment)
	e.GET("/payments/:id", h.GetPaymentByID)
	e.PUT("/payments/:id", h.UpdatePayment)
	e.DELETE("/payments/:id", h.DeletePayment)

	return e
}
