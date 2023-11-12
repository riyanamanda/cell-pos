package controllers

import (
	"cell-pos/server/config"
	"cell-pos/server/models"
	"cell-pos/server/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(ctx *gin.Context){
	var categories []models.Category

	config.DB.Find(&categories)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": categories,
	})
}

func StoreCategory(ctx *gin.Context){
	var input request.StoreCategoryRequest
	
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	category := models.Category{
		Name: input.Name,
	}
	config.DB.Create(&category)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data has been added successfully",
	})
}

func GetCategoryById(ctx *gin.Context){
	var category models.Category

	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes": false,
			"message": "Record not found",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": category,
	})
}

func UpdateCategoryById(ctx *gin.Context){
	var category models.Category

	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Record not found",
		})
	}

	var input request.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes": false,
			"message": err.Error(),
		})

		return
	}
	
	config.DB.Model(&category).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data has been updated successfully",
	})
}

func DeleteCategoryById(ctx *gin.Context){
	var category models.Category

	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	config.DB.Delete(&category)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data has been deleted successfully",
	})
}