package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(userID uuid.UUID) (string, error) {

	api_secret := os.Getenv("API_SECRET")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})

	// Replace "SomeAppSecret" with your actual secret key
	secretKey := []byte(api_secret)
	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

func ParseUserIDFromJWTToken(cookieWithJWT string) (string, error) {

	api_secret := os.Getenv("API_SECRET")
	hmacSecret := []byte(api_secret)

	token, err := jwt.Parse(cookieWithJWT, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена и другие параметры
		return hmacSecret, nil
	})

	if err != nil {
		// log.Printf("При извлечении токена произошла ошибка <%v>\n", err)
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Неверный JWT токен")
		return "", fmt.Errorf("invalid JWT Token ")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		log.Println("Не удалось извлечь USER_ID из токена")
		return "", fmt.Errorf("failed to parse claims")
	}

	return userID, nil
}
