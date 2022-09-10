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
	GetContactNum() (int, error)
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
	pageNum, err := strconv.Atoi(c.DefaultQuery(ginQueryPage, ginDefaultPageStart))
	if err != nil {
		return nil, err
	}

	return controller.service.FindAll(pageNum*retrieveResultLimit, retrieveResultLimit)
}

func (controller *contactController) Save(c *gin.Context) (int, error) {
	var newContact entities.ContactRequestBody

	if err := c.BindJSON(&newContact); err != nil {
		return 0, err
	}

	if err := ValidateContact(&newContact, true); err != nil {
		return 0, err
	}

	return controller.service.Save(&newContact)
}

func (controller *contactController) Delete(c *gin.Context) error {
	contactID, err := strconv.Atoi(c.Param(ginParamId))
	if err != nil {
		return err
	}

	controller.service.Delete(contactID)

	return nil
}

func (controller *contactController) Edit(c *gin.Context) error {
	contactID, err := strconv.Atoi(c.Param(ginParamId))
	if err != nil {
		return err
	}

	var updateContact entities.ContactRequestBody

	if err := c.BindJSON(&updateContact); err != nil {
		return err
	}

	if err := ValidateContact(&updateContact, false); err != nil {
		return err
	}

	return controller.service.Edit(&updateContact, contactID)
}

func (controller *contactController) GetContactNum() (int, error) {
	return controller.service.GetContactNum()
}

func (controller *contactController) Search(c *gin.Context) ([]entities.ContactResponseBody, error) {
	firstName, _ := c.GetQuery(ginQueryFirstName)

	lastName, _ := c.GetQuery(ginQueryLastName)

	pageNum, err := strconv.Atoi(c.DefaultQuery(ginQueryPage, ginDefaultPageStart))
	if err != nil {
		return nil, err
	}

	return controller.service.Search(firstName, lastName, pageNum)
}
