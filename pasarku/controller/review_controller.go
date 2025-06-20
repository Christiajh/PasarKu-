package controller

import (
	"net/http"
	"pasarku/helper"
	"pasarku/model"
	"pasarku/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateReview membuat ulasan untuk produk
func CreateReview(c *gin.Context) {
	var input structs.ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	userID := c.MustGet("userID").(uint)

	review := model.Review{
		UserID:    userID,
		ProductID: input.ProductID,
		Rating:    input.Rating,
		Comment:   input.Comment,
	}

	if err := model.CreateReview(&review); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to create review", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, "Review added", review)
}

// GetReviewsByProduct menampilkan semua review untuk 1 produk
func GetReviewsByProduct(c *gin.Context) {
	idStr := c.Param("product_id")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	reviews, err := model.GetReviewsByProductID(uint(productID))
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to get reviews", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Reviews found", reviews)
}

// UpdateReview mengedit komentar dan rating review
func UpdateReview(c *gin.Context) {
	var input structs.ReviewInput
	idStr := c.Param("id")
	reviewID, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid review ID", nil)
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	review, err := model.GetReviewByID(uint(reviewID))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Review not found", nil)
		return
	}

	userID := c.MustGet("userID").(uint)
	if review.UserID != userID {
		helper.ErrorResponse(c, http.StatusForbidden, "You can only update your own review", nil)
		return
	}

	review.Comment = input.Comment
	review.Rating = input.Rating

	if err := model.UpdateReview(&review); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to update review", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Review updated", review)
}

// DeleteReview menghapus review (oleh user atau admin)
func DeleteReview(c *gin.Context) {
	idStr := c.Param("id")
	reviewID, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid review ID", nil)
		return
	}

	review, err := model.GetReviewByID(uint(reviewID))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Review not found", nil)
		return
	}

	userID := c.MustGet("userID").(uint)
	userRole := c.MustGet("userRole").(string)

	if review.UserID != userID && userRole != "admin" {
		helper.ErrorResponse(c, http.StatusForbidden, "You are not allowed to delete this review", nil)
		return
	}

	if err := model.DeleteReview(&review); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete review", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Review deleted", nil)
}
