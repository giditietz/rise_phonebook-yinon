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
	addressService    service.AddressService       = service.NewAddressService()
	addressController controller.AddressController = controller.NewAddressController(addressService)
	phoneService      service.PhoneService         = service.NewPhoneService()
	phoneController   controller.PhoneController   = controller.NewPhoneController(phoneService)
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

func DeleteAddress(c *gin.Context) {
	err := addressController.Delete(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func DeletePhone(c *gin.Context) {
	err := phoneController.Delete(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}
