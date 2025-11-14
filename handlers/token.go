package handlers

import (
	"auth/internal/cryptography"
	"auth/internal/database/queries"
	"auth/internal/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RefreshToken(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := cryptography.ExtractToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		token, err := validation.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newToken, err := cryptography.GenerateRefreshToken(token.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = queries.RestoreRefreshToken(db, tokenString, newToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": newToken})
	}
}
