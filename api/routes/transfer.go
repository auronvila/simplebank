package routes

import "github.com/gin-gonic/gin"

func TransferRoutes(router gin.IRoutes, createTransfer gin.HandlerFunc) {
	router.POST("/transfer", createTransfer)

}
