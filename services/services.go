package services

import (
	daccess "jwt-gin-example/db"
	"jwt-gin-example/models"

	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type user = models.User

var secret = []byte("secret")
var userRepository = daccess.NewUserRepository()

type JwtService struct {
}

func (s *JwtService) GenerateJWT(User models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": models.User{
			ID:       User.ID,
			Name:     User.Name,
			Email:    User.Email,
			Password: User.Password,
			Role:     User.Role,
			FamilyId: User.FamilyId,
		},
		"exp": time.Now().Add(time.Minute * 5).Unix(),
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
		user := models.User{
			ID:       int64(userClaim["id"].(float64)),
			Name:     userClaim["name"].(string),
			Email:    userClaim["email"].(string),
			Password: userClaim["password"].(string),
			Role:     models.Role(userClaim["role"].(string)),
			FamilyId: int64(userClaim["familyId"].(float64)),
			Family:   models.Family{ID: int64(userClaim["familyId"].(float64))},
		}

		return VerifyResponse{User: user}, nil
	} else {
		return VerifyResponse{}, fmt.Errorf("invalid token")
	}
}

type UsersService struct {
}

func (s *UsersService) CreateUser(request models.RegisterRequest) (models.RegisterResponse, error) {
	user := user{Name: request.Name, Email: request.Email, Password: request.Password, Family: models.Family{Name: request.FamilyName}}

	user, err := userRepository.CreateUser(user)

	if err != nil {
		return models.RegisterResponse{}, err
	}

	return models.RegisterResponse{User: user}, err
}
