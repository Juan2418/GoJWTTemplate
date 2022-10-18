package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LoginRequest struct {
	Id int `json:"id"`
}
