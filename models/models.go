package models

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

func (ct *Role) Scan(value interface{}) error {
	*ct = Role(value.([]byte))
	return nil
}

func (ct Role) Value() (driver.Value, error) {
	return string(ct), nil
}

type User struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      Role      `json:"role" gorm:"default:'user':type:enum('admin', 'user')"`
	FamilyId  int64     `json:"familyId" gorm:"column:familyId"`
	Family    Family    `json:"family"`
}

type Family struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	ApiKey    string    `json:"apiKey" gorm:"column:apiKey;unique"`
	Members   []User    `json:"members"`
}

func (family *Family) BeforeCreate(tx *gorm.DB) (err error) {
	family.ApiKey = "family-costs-" + uuid.New().String()
	return
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return
}

func (user *User) AfterFind(tx *gorm.DB) (err error) {
	user.Password = ""
	return
}

func (user *User) ComparePassword(password string) bool {
	return CheckPasswordHash(password, user.Password)
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "User"
}

func (Family) TableName() string {
	return "Family"
}

type LoginRequest struct {
	Id int64 `json:"id"`
}

type RegisterRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FamilyName string `json:"familyName"`
}

type RegisterResponse struct {
	User User `json:"userId"`
}
