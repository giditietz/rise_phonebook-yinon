package server

import (
	"net/http"
	"phonebook/server/controller"
	"phonebook/server/service"

	"github.com/gin-gonic/gin"
)

var (
	contactService    service.ContactService       = service.NewContactService()
	contactController controller.ContactController = controller.NewContactController(contactService)
)

func GetAllContacts(c *gin.Context) {
	contacts, err := contactController.FindAll(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, contacts)
}

func CreateContact(c *gin.Context) {
	contactID, err := contactController.Save(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, contactID)
}

func DeleteContact(c *gin.Context) {
	err := contactController.Delete(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func EditContact(c *gin.Context) {
	err := contactController.Edit(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func SearchContact(c *gin.Context) {
	contacts, err := contactController.Search(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, contacts)
}

func GetNumOfContact(c *gin.Context) {
	contactsNum, err := contactController.GetContactNum()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, contactsNum)
}
