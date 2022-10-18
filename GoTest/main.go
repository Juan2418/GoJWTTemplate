package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users []user = []user{
	{ID: 1, Name: "user1"},
	{ID: 2, Name: "user2"},
	{ID: 3, Name: "user3"},
}

func (usr user) toString() string {
	return `{"id":` + string(usr.ID) + `,"name":"` + usr.Name + `"}`
}

func generateJWT(User user) (string, error) {
	var secret = []byte("secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": User,
		"exp":  time.Now().Add(time.Hour * 24),
	})

	return token.SignedString(secret)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
}

type LoginRequest struct {
	Id int `json:"id"`
}

func login(c *gin.Context) {
	// bodyId := c.PostForm("id")
	// userId, _ := strconv.Atoi(bodyId)
	// user := users[userId]
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

	token, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "requestedId": requestBody.Id, "user": user})
}

func main() {
	u := user{ID: 1, Name: "John"}

	r := gin.Default()
	r.GET("/user/:id/:name", getUser)
	r.POST("/login", login)
	r.Run(":8080")

	println(u.Name)
}
