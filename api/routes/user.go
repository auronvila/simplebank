package routes

import "github.com/gin-gonic/gin"

func UserRoutes(router gin.IRoutes, createUser, loginUser gin.HandlerFunc) {
	router.POST("/user", createUser)
	router.POST("/user/login", loginUser)
}
