package delivery

import (
	"net/http"
	"strconv"

	"miniproject_go_wardahfdn/app/model"
	"miniproject_go_wardahfdn/app/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	UserUseCase       usecase.UserUseCase
	RestaurantUseCase usecase.RestaurantUseCase
	MenuUseCase       usecase.MenuUseCase
	OrderUseCase      usecase.OrderUseCase
	PaymentUseCase    usecase.PaymentUseCase
}

func NewHandler(userUC usecase.UserUseCase, restaurantUC usecase.RestaurantUseCase,
	menuUC usecase.MenuUseCase, orderUC usecase.OrderUseCase, paymentUC usecase.PaymentUseCase) *Handler {
	return &Handler{
		UserUseCase:       userUC,
		RestaurantUseCase: restaurantUC,
		MenuUseCase:       menuUC,
		OrderUseCase:      orderUC,
		PaymentUseCase:    paymentUC,
	}
}

// User handlers
func (h *Handler) CreateUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	createdUser := h.UserUseCase.CreateUser(&user)
	if createdUser == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
	}
	return c.JSON(http.StatusOK, createdUser)
}

func (h *Handler) GetAllUser(c echo.Context) error {
	users, err := h.UserUseCase.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserByID(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	user, err := h.UserUseCase.GetUserByID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err = h.UserUseCase.UpdateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	err = h.UserUseCase.DeleteUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

// Restaurant handlers
func (h *Handler) CreateRestaurant(c echo.Context) error {
	var restaurant model.Restaurant
	err := c.Bind(&restaurant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	createdRestaurant := h.RestaurantUseCase.CreateRestaurant(&restaurant)
	if createdRestaurant == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create restaurant"})
	}
	return c.JSON(http.StatusOK, createdRestaurant)
}

func (h *Handler) GetAllRestaurant(c echo.Context) error {
	restaurant, err := h.RestaurantUseCase.GetAllRestaurant()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, restaurant)
}

func (h *Handler) GetRestaurantByID(c echo.Context) error {
	restaurantIDStr := c.Param("id")
	restaurantID, err := strconv.Atoi(restaurantIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid restaurant ID"})
	}
	restaurant, err := h.RestaurantUseCase.GetRestaurantByID(uint(restaurantID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, restaurant)
}

func (h *Handler) UpdateRestaurant(c echo.Context) error {
	var restaurant model.Restaurant
	err := c.Bind(&restaurant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err = h.RestaurantUseCase.UpdateRestaurant(&restaurant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, restaurant)
}

func (h *Handler) DeleteRestaurant(c echo.Context) error {
	restaurantID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid restaurant ID"})
	}

	err = h.RestaurantUseCase.DeleteRestaurant(uint(restaurantID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Restaurant deleted successfully"})
}

// Menu handlers
func (h *Handler) CreateMenu(c echo.Context) error {
	var menu model.Menu
	err := c.Bind(&menu)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	createdMenu := h.MenuUseCase.CreateMenu(&menu)
	if createdMenu == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create menu"})
	}
	return c.JSON(http.StatusOK, createdMenu)
}

func (h *Handler) GetAllMenu(c echo.Context) error {
	menu, err := h.MenuUseCase.GetAllMenu()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, menu)
}

func (h *Handler) GetMenuByID(c echo.Context) error {
	menuIDStr := c.Param("id")
	menuID, err := strconv.Atoi(menuIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid menu ID"})
	}
	menu, err := h.MenuUseCase.GetMenuByID(uint(menuID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, menu)
}

func (h *Handler) UpdateMenu(c echo.Context) error {
	var menu model.Menu
	err := c.Bind(&menu)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err = h.MenuUseCase.UpdateMenu(&menu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, menu)
}

func (h *Handler) DeleteMenu(c echo.Context) error {
	menuID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid menu ID"})
	}

	err = h.MenuUseCase.DeleteMenu(uint(menuID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Menu deleted successfully"})
}

// Order handlers
func (h *Handler) CreateOrder(c echo.Context) error {
	var order model.Order
	err := c.Bind(&order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	createdOrder := h.OrderUseCase.CreateOrder(&order)
	if createdOrder == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create order"})
	}
	return c.JSON(http.StatusOK, createdOrder)
}

func (h *Handler) GetAllOrder(c echo.Context) error {
	order, err := h.OrderUseCase.GetAllOrder()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, order)
}

func (h *Handler) GetOrderByID(c echo.Context) error {
	orderIDStr := c.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}
	order, err := h.OrderUseCase.GetOrderByID(uint(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, order)
}

func (h *Handler) UpdateOrder(c echo.Context) error {
	var order model.Order
	err := c.Bind(&order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err = h.OrderUseCase.UpdateOrder(&order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, order)
}

func (h *Handler) DeleteOrder(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}

	err = h.OrderUseCase.DeleteOrder(uint(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Order deleted successfully"})
}

// Payment handlers
func (h *Handler) CreatePayment(c echo.Context) error {
	var payment model.Payment
	err := c.Bind(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	createdPayment := h.PaymentUseCase.CreatePayment(&payment)
	if createdPayment == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create payment"})
	}
	return c.JSON(http.StatusOK, createdPayment)
}

func (h *Handler) GetAllPayment(c echo.Context) error {
	payment, err := h.PaymentUseCase.GetAllPayment()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, payment)
}

func (h *Handler) GetPaymentByID(c echo.Context) error {
	paymentIDStr := c.Param("id")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payment ID"})
	}
	payment, err := h.PaymentUseCase.GetPaymentByID(uint(paymentID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, payment)
}

func (h *Handler) UpdatePayment(c echo.Context) error {
	var payment model.Payment
	err := c.Bind(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err = h.PaymentUseCase.UpdatePayment(&payment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, payment)
}

func (h *Handler) DeletePayment(c echo.Context) error {
	paymentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payment ID"})
	}

	err = h.PaymentUseCase.DeletePayment(uint(paymentID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Payment deleted successfully"})
}
