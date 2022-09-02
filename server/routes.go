package server

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/contacts", GetAllContacts)
		api.POST("/contacts", CreateContact)
		api.DELETE("/contacts/:id", DeleteContact)
	}

	router.Run("localhost:9000")
}
