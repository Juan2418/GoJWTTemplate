package daccess

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "user:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
var Client, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// type UserRepository interface {
// 	FindUserById(id int) (*db.UserModel, error)
// }

// type userRepository struct {
// }

// func NewUserRepository() UserRepository {
// 	return &userRepository{}
// }

// func (r *userRepository) FindUserById(id int) (*db.UserModel, error) {
// 	return Client.User.FindUnique(db.User.ID.Equals(id)).Exec(context.Background())
// }
