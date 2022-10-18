package main

import (
	"jwt-gin-example/models"
	"jwt-gin-example/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type user = models.User
var JWTService = services.JwtService{}

var users []user = []user{
	{ID: 1, Name: "user1"},
	{ID: 2, Name: "user2"},
	{ID: 3, Name: "user3"},
}

func GetVerifiedToken(c *gin.Context) {
	tokenString := GetToken(c.Request.Header)
	response, err := JWTService.VerifyToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": response.User})
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

	token, err := JWTService.GenerateJWT(user)
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
