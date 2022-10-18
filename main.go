package main

import (
	"jwt-gin-example/models"

	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type user = models.User

var users []user = []user{
	{ID: 1, Name: "user1"},
	{ID: 2, Name: "user2"},
	{ID: 3, Name: "user3"},
}

var secret = []byte("secret")

func GenerateJWT(User user) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": User,
		"exp":  time.Now().Add(time.Minute * 5).Unix(),
	})

	return token.SignedString(secret)
}

func GetVerifiedToken(c *gin.Context) {
	tokenString := GetToken(c.Request.Header)
	response, err := VerifyToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": response.User})
}

type VerifyResponse struct {
	User user `json:"user"`
}

func VerifyToken(tokenString string) (VerifyResponse, error) {

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

func GetToken(Header http.Header) string {
	tokenString := Header.Get("Authorization")
	tokenString = tokenString[7:]
	return tokenString
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
}

type LoginRequest struct {
	Id int `json:"id"`
}

func Login(c *gin.Context) {
	var requestBody LoginRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestBody.Id >= len(users) || requestBody.Id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id", "id": requestBody.Id})
		return
	}

	user := users[requestBody.Id]

	token, err := GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "requestedId": requestBody.Id, "user": user})
}

func main() {
	u := user{ID: 1, Name: "John"}

	r := gin.Default()
	r.GET("/user/:id/:name", GetUser)
	r.POST("/login", Login)
	r.GET("/verify", GetVerifiedToken)
	r.Run(":8080")

	println(u.Name)
}
