package handler

import (
	"app/app/model"
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser model.User

		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
			return
		}

		CreateNewUser(c, db, newUser)
	}
}

func SignIn(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed signin"})
			return
		}
		findUser, err := GetUser(db, user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found user"})
			return
		}

		user.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password)))
		if findUser.Email != user.Email || findUser.Password != user.Password {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
			return
		}
		fmt.Println(findUser.ID)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": findUser.ID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte("MY_TODO_APP_SECRET_KEY"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error signing token"})
			return
		}

		c.Header("Authorization", tokenString)
		c.JSON(http.StatusOK, tokenString)
	}
}

func ParseJWT(tokenStr string) (*model.Claims, error) {
	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("MY_TODO_APP_SECRET_KEY"), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
