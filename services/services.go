package services

import (
	"jwt-gin-example/models"

	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type user = models.User

var secret = []byte("secret")

type JwtService struct {
}

func (s *JwtService) GenerateJWT(User user) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": User,
		"exp":  time.Now().Add(time.Minute * 5).Unix(),
	})

	return token.SignedString(secret)
}

type VerifyResponse struct {
	User user `json:"user"`
}

func (s *JwtService) VerifyToken(tokenString string) (VerifyResponse, error) {

	println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return secret, nil
	})

	if err != nil {
		return VerifyResponse{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userClaim := claims["user"].(map[string]interface{})
		id := int(userClaim["id"].(float64))
		name := userClaim["name"].(string)
		return VerifyResponse{User: user{ID: id, Name: name}}, nil
	} else {
		return VerifyResponse{}, fmt.Errorf("invalid token")
	}
}
