package controller

import (
	"net/http"
	"project/database/model"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

}

func FindInDatabase(c *gin.Context) {
	var user model.Person
	var foundUser model.Person

	c.BindJSON(&user)

	if user.IsValidEmail() || repository.IsValidPassword(c, person, foundPerson) {
		return
	}

	repository.Update(ctx, foundPerson)

	c.JSON(http.StatusOK, &foundPerson)
}
