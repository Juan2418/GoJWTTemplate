package models

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:true" gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	FamilyId  int64     `json:"familyId"`
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "User"
}

type LoginRequest struct {
	Id int `json:"id"`
}
