package server

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()
	router.GET("/api/contacts", GetAllContacts)
	router.POST("/api/contacts", CreateContact)

	router.Run("localhost:9000")
}
