package validation

import (
	"auth/internal/cryptography"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateAccessToken(accessToken string) (*cryptography.Token, error) {
	secretKey, err := cryptography.GetJWTSecretKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неподдерживаемый метод подписи")
		}
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("недействительный или истекший токен")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("не удалось извлечь claims из access токена")
	}

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("не удалось извлечь id из claims")
	}

	expiresFloat, ok := claims["expires_at"].(float64)
	if !ok {
		return nil, errors.New("не удалось извлечь expires_at из claims")
	}

	tokenStruct := &cryptography.Token{
		ID:        int(idFloat),
		ExpiresAt: time.Unix(int64(expiresFloat), 0),
	}

	if time.Now().After(tokenStruct.ExpiresAt) {
		return nil, errors.New("токен истек")
	}

	return tokenStruct, nil
}
