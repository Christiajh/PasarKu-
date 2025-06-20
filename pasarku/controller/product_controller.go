package controller

import (
	"net/http"
	"pasarku/helper"
	"pasarku/model"
	"pasarku/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
func CreateProduct(c *gin.Context) {
	var input structs.ProductInput

	// Validasi input
	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	// Ambil user ID dari JWT middleware
	userID := c.MustGet("userID").(uint)

	product := model.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Category:    input.Category,
		Tag:         input.Tag,
		UserID:      userID,
	}

	if err := model.CreateProduct(&product); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to create product", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, "Product created", product)
}

// GetAllProducts godoc
func GetAllProducts(c *gin.Context) {
	products, err := model.GetAllProducts()
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch products", err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "All products", products)
}

// GetProductByID godoc
func GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	product, err := model.GetProductByID(uint(id))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Product not found", nil)
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "Product found", product)
}

// UpdateProduct godoc
func UpdateProduct(c *gin.Context) {
	var input structs.ProductInput
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	product, err := model.GetProductByID(uint(id))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Product not found", nil)
		return
	}

	// Cek hak akses
	userID := c.MustGet("userID").(uint)
	userRole := c.MustGet("userRole").(string)
	if product.UserID != userID && userRole != "admin" {
		helper.ErrorResponse(c, http.StatusForbidden, "You are not allowed to update this product", nil)
		return
	}

	// Update data
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Category = input.Category
	product.Tag = input.Tag

	if err := model.UpdateProduct(&product); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to update product", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Product updated", product)
}

// DeleteProduct godoc
func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	product, err := model.GetProductByID(uint(id))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Product not found", nil)
		return
	}

	userID := c.MustGet("userID").(uint)
	userRole := c.MustGet("userRole").(string)
	if product.UserID != userID && userRole != "admin" {
		helper.ErrorResponse(c, http.StatusForbidden, "You are not allowed to delete this product", nil)
		return
	}

	if err := model.DeleteProduct(&product); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete product", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Product deleted", nil)
}

// UploadProductImage godoc
func UploadProductImage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	product, err := model.GetProductByID(uint(id))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Product not found", nil)
		return
	}

	userID := c.MustGet("userID").(uint)
	userRole := c.MustGet("userRole").(string)
	if product.UserID != userID && userRole != "admin" {
		helper.ErrorResponse(c, http.StatusForbidden, "You are not allowed to upload image for this product", nil)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Image file is required", nil)
		return
	}

	
	path := "public/uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to upload image", err.Error())
		return
	}

	product.ImageURL = "/" + path
	if err := model.UpdateProduct(&product); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to save image URL", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Image uploaded successfully", product)
}
