package daccess

import (
	"jwt-gin-example/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "user:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
var Client, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type UserRepository interface {
	FindUserById(id int) (models.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindUserById(id int) (models.User, error) {
	var user models.User
	resultContext := Client.First(user, id)

	return user, resultContext.Error
}
