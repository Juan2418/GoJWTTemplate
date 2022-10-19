package main

import (
	daccess "jwt-gin-example/db"
	"jwt-gin-example/models"
	"jwt-gin-example/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type user = models.User

var JWTService = services.JwtService{}

var usersRepository = daccess.NewUserRepository()

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

func Login(c *gin.Context) {
	var requestBody models.LoginRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := usersRepository.FindUserById(requestBody.Id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := JWTService.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "requestedId": requestBody.Id, "user": user})
}

func PingDB(c *gin.Context) {
	var user models.User
	daccess.Client.First(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func main() {
	r := gin.Default()
	r.POST("/login", Login)
	r.GET("/verify", GetVerifiedToken)
	r.GET("/ping", PingDB)
	r.Run(":8080")
}
