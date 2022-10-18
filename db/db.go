package daccess

import "jwt-gin-example/models"

type UserRepository interface {
	FindUserById(id int) (models.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindUserById(id int) (models.User, error) {
	return models.User{ID: id, Name: "user" + string(id)}, nil
}
