package handler

import (
	"app/app/model"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNewUser(c *gin.Context, db *gorm.DB, newUser model.User) {
	newUser.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(newUser.Password)))
	res := db.Table("users").Create(&newUser)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed signUp"})
		return
	}
}

var findUser model.User

func GetUser(db *gorm.DB, user model.User) (*model.User, error) {
	email := user.Email
	findUser.ID = 0
	if err := db.Table("users").Where("email = ?", email).First(&findUser).Error; err != nil {
		fmt.Println("can not found user")
		return nil, err
	} else {
		fmt.Println(findUser)
		return &findUser, nil
	}
}

func GetCurrentUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(int)
		var user model.User
		if err := db.Table("users").Where("id = ?", userID).First(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
