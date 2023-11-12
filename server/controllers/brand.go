package controllers

import (
	"cell-pos/server/config"
	"cell-pos/server/models"
	"cell-pos/server/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBrands(ctx *gin.Context) {
	var brands []models.Brand
	config.DB.Find(&brands)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": brands,
	})
}

func StoreBrand(ctx *gin.Context){
	var input request.StoreBrandRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
	}

	brand := models.Brand{
		Name: input.Name,
	}
 	config.DB.Create(&brand)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data have been added successfully",
	})
}

func GetBrandById(ctx *gin.Context){
	var brand models.Brand

	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&brand).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Record not found",
		})

		return
	  }

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": brand,
	})
}

func UpdateBrandById(ctx *gin.Context){
	var brand models.Brand

	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&brand).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Record not found",
		})
	}

	var input request.UpdateBrandRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})

		return
	}

	config.DB.Model(&brand).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data has been updated successfully",
	})
}

func DeleteBrand(ctx *gin.Context){
	var brand models.Brand

	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&brand).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Record not found",
		})

		return
	}

	config.DB.Delete(&brand)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data has been deleted successfully",
	})
}