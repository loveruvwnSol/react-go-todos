package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Base
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
