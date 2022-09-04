package controller

import (
	"phonebook/server/entities"
	"phonebook/server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactController interface {
	FindAll(c *gin.Context) ([]entities.ContactResponseBody, error)
	Search(c *gin.Context) ([]entities.ContactResponseBody, error)
	Save(c *gin.Context) (int, error)
	Delete(c *gin.Context) error
	Edit(c *gin.Context) error
}

type contactController struct {
	service service.ContactService
}

func NewContactController(service service.ContactService) ContactController {
	return &contactController{
		service: service,
	}
}

func (controller *contactController) FindAll(c *gin.Context) ([]entities.ContactResponseBody, error) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery(ginQueryPage, ginDefaultPageStart))

	ret, err := controller.service.FindAll(pageNum*retrieveResultLimit, retrieveResultLimit)

	return ret, err
}

func (controller *contactController) Save(c *gin.Context) (int, error) {
	var newContact entities.ContactRequestBody

	if err := c.BindJSON(&newContact); err != nil {
		return 0, err
	}
	ret, err := controller.service.Save(&newContact)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func (controller *contactController) Delete(c *gin.Context) error {
	contactID, _ := strconv.Atoi(c.Param(ginParamId))

	controller.service.Delete(contactID)

	return nil
}

func (controller *contactController) Edit(c *gin.Context) error {
	contactID, _ := strconv.Atoi(c.Param(ginParamId))

	var updateContact entities.ContactRequestBody

	if err := c.BindJSON(&updateContact); err != nil {
		return err
	}

	return controller.service.Edit(&updateContact, contactID)
}

func (controller *contactController) Search(c *gin.Context) ([]entities.ContactResponseBody, error) {
	firstName, _ := c.GetQuery(ginQueryFirstName)
	lastName, _ := c.GetQuery(ginQueryLastName)
	pageNum, _ := strconv.Atoi(c.DefaultQuery(ginQueryPage, ginDefaultPageStart))

	return controller.service.Search(firstName, lastName, pageNum)
}
