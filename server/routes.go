package server

import (
	"phonebook/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSMiddleware())

	api := server.Group("/api")
	{
		api.GET("/contacts", GetAllContacts)
		api.POST("/contacts", CreateContact)
		api.DELETE("/contacts/:id", DeleteContact)
		api.PUT("/contacts/:id", EditContact)
		api.GET("/contacts/search", SearchContact)
	}

	go server.Run("0.0.0.0:9000")
	second_server := gin.New()

	second_server.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSMiddleware())

	second_api := second_server.Group("/api")
	{
		second_api.GET("/contacts", GetAllContacts)
		second_api.POST("/contacts", CreateContact)
		second_api.DELETE("/contacts/:id", DeleteContact)
		second_api.PUT("/contacts/:id", EditContact)
		second_api.GET("/contacts/search", SearchContact)
	}
	second_server.Run("0.0.0.0:9001")
}
