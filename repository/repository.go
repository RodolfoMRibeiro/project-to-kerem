package repository

import (
	"context"
	"net/http"
)

func IsValidEmail(c *gin.Context, ctx context.Context, person models.Person, foundPerson *models.Person) bool {
	if err := database.Collection(database.Database(), models.TABLE).
		FindOne(ctx, bson.M{"email": person.Email}).Decode(&foundPerson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "email is not correct"})
		return true
	}

	return false
}
