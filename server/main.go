package main

import (
	"cell-pos/server/config"
	"cell-pos/server/controllers"
	"cell-pos/server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	config.Connection()

	router.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{"message" : "Welcome to Cell POS API"})
	})
	api := router.Group("/api")
	{
		api.GET("/dashboard", controllers.Dashboard)
	}

	server := http.Server{
		Addr: ":8000",
		Handler: router,
	}

	err := server.ListenAndServe();
	helper.PanicIfError(err)
}