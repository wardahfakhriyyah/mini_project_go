package delivery

import (
	"net/http"

	"miniproject_go_wardahfdn/app/model"
	"miniproject_go_wardahfdn/app/usecase"

	"github.com/labstack/echo/v4"
)

type HttpDelivery struct {
	UserUsecase       usecase.UserUsecase
	RestaurantUsecase usecase.RestaurantUsecase
	MenuUsecase       usecase.MenuUsecase
	OrderUsecase      usecase.OrderUsecase
	PaymentUsecase    usecase.PaymentUsecase
}

func NewHttpDelivery(e *echo.Echo, userUsecase usecase.UserUsecase, restaurantUsecase usecase.RestaurantUsecase,
	menuUsecase usecase.MenuUsecase, orderUsecase usecase.OrderUsecase, paymentUsecase usecase.PaymentUsecase) {
	handler := HttpDelivery{
		UserUsecase:       userUsecase,
		RestaurantUsecase: restaurantUsecase,
		MenuUsecase:       menuUsecase,
		OrderUsecase:      orderUsecase,
		PaymentUsecase:    paymentUsecase,
	}

	e.GET("/", handler.Homepage)
	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)

	user := e.Group("/users")
	user.GET("", handler.GetAllUsers)
	user.GET("/:id", handler.GetUserById)
	user.PUT("/:id", handler.UpdateUser)
	user.DELETE("/:id", handler.DeleteUser)

	restaurant := e.Group("/restaurants")
	restaurant.GET("", handler.GetAllRestaurants)
	restaurant.GET("/:id", handler.GetRestaurantById)
	restaurant.POST("", handler.CreateRestaurant)
	restaurant.PUT("/:id", handler.UpdateRestaurant)
	restaurant.DELETE("/:id", handler.DeleteRestaurant)

	menu := e.Group("/menus")
	menu.GET("", handler.GetAllMenus)
	menu.GET("/:id", handler.GetMenuById)
	menu.POST("", handler.CreateMenu)
	menu.PUT("/:id", handler.UpdateMenu)
	menu.DELETE("/:id", handler.DeleteMenu)

	order := e.Group("/orders")
	order.GET("", handler.GetAllOrders)
	order.GET("/:id", handler.GetOrderById)
	order.POST("", handler.CreateOrder)
	order.PUT("/:id", handler.UpdateOrder)
	order.DELETE("/:id", handler.DeleteOrder)

	payment := e.Group("/payments")
	payment.GET("", handler.GetAllPayments)
	payment.GET("/:id", handler.GetPaymentById)
	payment.POST("", handler.CreatePayment)
	payment.PUT("/:id", handler.UpdatePayment)
	payment.DELETE("/:id", handler.DeletePayment)
}

func (h *HttpDelivery) Homepage(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Mini Project Wardah Foundation")
}

func (h *HttpDelivery) Register(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	err := h.UserUsecase.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func (h *HttpDelivery) Login(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
}
