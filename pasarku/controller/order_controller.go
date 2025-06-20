package controller

import (
	"net/http"
	"pasarku/helper"
	"pasarku/model"
	"pasarku/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrder membuat pesanan baru
func CreateOrder(c *gin.Context) {
	var input structs.OrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	userID := c.MustGet("userID").(uint)

	order := model.Order{
		UserID:    userID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Status:    "pending",
	}

	if err := model.CreateOrder(&order); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to create order", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, "Order created", order)
}

// GetMyOrders menampilkan semua pesanan milik user login
func GetMyOrders(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	orders, err := model.GetOrdersByUserID(userID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to get orders", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Orders retrieved", orders)
}

// GetAllOrders hanya untuk admin melihat semua pesanan
func GetAllOrders(c *gin.Context) {
	userRole := c.MustGet("userRole").(string)
	if userRole != "admin" {
		helper.ErrorResponse(c, http.StatusForbidden, "Only admin can view all orders", nil)
		return
	}

	orders, err := model.GetAllOrders()
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to get all orders", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "All orders", orders)
}

// GetOrderByID menampilkan detail satu pesanan
func GetOrderByID(c *gin.Context) {
	idStr := c.Param("id")
	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid order ID", nil)
		return
	}

	order, err := model.GetOrderByID(uint(orderID))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Order not found", nil)
		return
	}

	userID := c.MustGet("userID").(uint)
	userRole := c.MustGet("userRole").(string)

	if order.UserID != userID && userRole != "admin" && userRole != "seller" {
		helper.ErrorResponse(c, http.StatusForbidden, "Access denied", nil)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Order detail", order)
}

// UpdateOrderStatus mengubah status pesanan → hanya admin/seller
func UpdateOrderStatus(c *gin.Context) {
	idStr := c.Param("id")
	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid order ID", nil)
		return
	}

	var input structs.OrderStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid status input", err.Error())
		return
	}

	order, err := model.GetOrderByID(uint(orderID))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Order not found", nil)
		return
	}

	userRole := c.MustGet("userRole").(string)
	if userRole != "admin" && userRole != "seller" {
		helper.ErrorResponse(c, http.StatusForbidden, "You are not allowed to update order status", nil)
		return
	}

	order.Status = input.Status
	if err := model.UpdateOrder(&order); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to update order", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Order status updated", order)
}
