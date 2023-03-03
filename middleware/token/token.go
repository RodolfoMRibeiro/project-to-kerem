package token

import (
	"project/commom/env"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	FirstName string
	LastName  string
	Email     string
	UserType  string
	Uid       string
	jwt.StandardClaims
}

func GenerateToken(firstName string, lastName string, email string, userType string, uid string) (signedToken string) {
	expirationTimeOfOneDay := time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()

	claims := &SignedDetails{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UserType:  userType,
		Uid:       uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTimeOfOneDay,
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(env.SECRET_KEY))

	return token
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, _ := jwt.ParseWithClaims(
		signedToken,
		claims,
		keyFunction,
	)

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "token is invalid"
		return
	}

	if localTime := time.Now().Local().Unix(); claims.ExpiresAt < localTime {
		msg = "token is expired"
		return
	}

	return claims, msg
}

func keyFunction(token *jwt.Token) (interface{}, error) {
	return []byte(env.SECRET_KEY), nil
}
