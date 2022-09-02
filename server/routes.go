package server

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()
	router.GET("/api/contacts", GetAllContacts)
	router.GET("/api/contacts/:id", GetContact)

	router.Run("localhost:9000")
}
