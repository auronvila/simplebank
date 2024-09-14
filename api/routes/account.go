package routes

import "github.com/gin-gonic/gin"

func AccountRoutes(router *gin.Engine, createAccount, getAccount, listAccount, updateAccount, deleteAccount gin.HandlerFunc) {
	router.POST("/accounts", createAccount)
	router.GET("/account/:id", getAccount)
	router.GET("/accounts", listAccount)
	router.PUT("/account/:id", updateAccount)
	router.DELETE("/account/:id", deleteAccount)
}
