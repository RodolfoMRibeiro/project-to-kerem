package controller

import (
	"jwt-project/dto"
	"jwt-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{service: service}
}

func (control controller) SignUp(c *gin.Context) {
	var dtoPerson dto.DtoSignUp
	c.BindJSON(&dtoPerson)

	insert, err := control.service.InsertInDatabase(c, dtoPerson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, insert)
}

func (control controller) Login(c *gin.Context) {
	var dtoPerson dto.DtoLogIn
	c.BindJSON(&dtoPerson)

	foundPerson, err := control.service.FindInDatabase(c, dtoPerson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &foundPerson.ID)
}

func (control controller) GetUser(c *gin.Context) {
	var dtoPerson dto.GetUser

	personId := c.Param("userid")

	person, err := control.service.GetFromDatabase(c, dtoPerson, personId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (control controller) GetUsers(c *gin.Context) {
	var allUsers []primitive.M

	allUsersList, err := control.service.GetallFromDatabase(c, allUsers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, allUsersList)
}
