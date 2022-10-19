package daccess

import (
	"fmt"
	"jwt-gin-example/models"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "user:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
var Client, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type UserRepository interface {
	FindUserById(id int64) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindUserById(id int64) (models.User, error) {
	var user models.User
	resultContext := Client.First(&user, id)

	return user, resultContext.Error
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	resultContext := Client.Create(&user)

	user.Password = ""
	err := resultContext.Error

	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			return user, fmt.Errorf("email is already taken")
		}

		if strings.Contains(err.Error(), "familyId") {
			return user, fmt.Errorf("family already exists")
		}

		return user, resultContext.Error
	}

	return user, resultContext.Error
}
