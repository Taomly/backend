package handlers

import (
	"auth/internal/cryptography"
	"auth/internal/database/queries"
	"auth/internal/validation"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Signup(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Username string `json:"username" binding:"required"`
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
			return
		}

		err := validation.ValidateUserPassword(request.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hash, err := cryptography.HashPassword(request.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = queries.CreateUser(db, request.Username, request.Email, hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err != nil {
			log.Println(err)
			return
		}

		c.JSON(200, gin.H{
			"message": "user created",
		})
	}
}

func Login(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
			return
		}

		user, err := queries.GetUser(db, request.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		check := cryptography.VerifyPassword(request.Password, user.Password)
		if !check {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid login or password"})
			return
		}

		accessToken, err := cryptography.GenerateAccessToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		refreshToken, err := cryptography.GenerateRefreshToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = queries.StoreRefreshToken(db, refreshToken, time.Now().Add(time.Hour*24*30).Unix())

		c.JSON(200, gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		})
	}
}

func ResetPassword(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Password    string `json:"password" binding:"required"`
			NewPassword string `json:"new_password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
			return
		}

		if request.Password == request.NewPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "new password cannot be equal to old password"})
			return
		}

		err := validation.ValidateUserPassword(request.NewPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tokenString, err := cryptography.ExtractToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		token, err := validation.ValidateAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := queries.GetUserByID(db, token.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		check := cryptography.VerifyPassword(request.Password, user.Password)
		if !check {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid login or password"})
			return
		}

		newPassword, err := cryptography.HashPassword(request.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = queries.UpdateUserPassword(db, user.ID, newPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "password updated",
		})
	}
}
