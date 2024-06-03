package utils

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func ParseUserIDFromJWTToken(cookieWithJWT string) (string, error) {

	hmacSecretString := "SomeAppSecret"
	hmacSecret := []byte(hmacSecretString)

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
