package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Age  uint32
}

type UserRepository interface {
	Get(id uint) (*User, error)
}
