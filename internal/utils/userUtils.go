package utils

import (
	"log"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
	"github.com/google/uuid"
)

// Извлекаем статью по ID
func GetUserByIDFromDB(userUUID string) models.User {

	db := database.DB
	var user models.User // article slice

	// Извлекаем статью вместе с автором и категориями
	result := db.First(&user, "ID = ?", userUUID)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve user: " + result.Error.Error())
	}

	if user.ID == uuid.Nil {
		// handle error
		panic("failed to retrieve user: " + result.Error.Error())
	}

	log.Println("Пользователь — успешно извлечен:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.Name)
	log.Printf("	Е-mail: <%s>\n", user.Email)
	// log.Printf("	Статьи: <%v>\n", user.Articles)

	return user
}
