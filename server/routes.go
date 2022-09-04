package server

import (
	"phonebook/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	api := server.Group("/api")
	{
		api.GET("/contacts", GetAllContacts)
		api.POST("/contacts", CreateContact)
		api.DELETE("/contacts/:id", DeleteContact)
		api.PUT("/contacts/:id", EditContact)
		api.GET("/contacts/search", SearchContact)
	}

	server.Run("localhost:9000")
}
