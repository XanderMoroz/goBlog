package handlers

import (
	"log"
	// "net/http"

	// "github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
)

// func CreateUserInDB() {

// 	db := database.DB
// 	newUser := models.User{
// 		FirstName: "Jane",
// 		LastName:  "Doe",
// 		Email:     "jane111doe@gmail.com",
// 		Country:   "Spain",
// 		Role:      "Chef",
// 		Age:       30,
// 	}

// 	// Add a uuid to a new user record...
// 	newUser.ID = uuid.New()
// 	// ... Create a new user record...
// 	result := db.Create(&newUser)
// 	if result.Error != nil {
// 		panic("failed to create user: " + result.Error.Error())
// 	}
// 	// ... Handle successful creation ...
// 	log.Println("Новый пользователь — успешно создан:")
// 	log.Printf("	ID: <%s>\n", newUser.ID)
// 	log.Printf("	Имя: <%s>\n", newUser.FirstName)
// }

// func GetUsersFromDB() {

// 	db := database.DB
// 	var users []models.User // users slice

// 	result := db.Find(&users)

// 	if result.Error != nil {
// 		// handle error
// 		panic("failed to retrieve users: " + result.Error.Error())
// 	}

// 	log.Println("Список пользователей — успешно извлечен:")
// 	for _, user := range users {
// 		log.Printf("User ID: %s, Name: %s %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
// 	}

// }

// func GetUserByIdFromDB() {

// 	db := database.DB
// 	var users []models.User // users slice

// 	result := db.Where("ID = ?", "073b515e-f9b6-45b2-aeb1-bbaa6de48286").Find(&users)

// 	if result.Error != nil {
// 		panic("failed to retrieve user: " + result.Error.Error())
// 	}
// 	// iterate over the users slice and print the details of each user
// 	log.Println("Пользователь — успешно извлечен:")
// 	for _, user := range users {
// 		log.Printf("	ID: <%s>\n", user.ID)
// 		log.Printf("	Имя: <%s>\n", user.FirstName)
// 		log.Printf("	Фамилия: <%s>\n", user.LastName)
// 		log.Printf("	E-mail: <%s>\n", user.Email)
// 	}

// }

// func GetUserByNameFromDB() {

// 	db := database.DB
// 	var users []models.User // users slice

// 	result := db.Where("First_Name = ?", "Jane").Find(&users)

// 	if result.Error != nil {
// 		panic("failed to retrieve user: " + result.Error.Error())
// 	}
// 	// iterate over the users slice and print the details of each user
// 	log.Println("Пользователь — успешно извлечен:")
// 	for _, user := range users {
// 		log.Printf("	ID: <%s>\n", user.ID)
// 		log.Printf("	Имя: <%s>\n", user.FirstName)
// 		log.Printf("	Фамилия: <%s>\n", user.LastName)
// 		log.Printf("	E-mail: <%s>\n", user.Email)
// 	}

// }

func UpdateUserByIdFromDB() {

	db := database.DB

	var user models.User

	// Retrieve the record you want to update
	result := db.First(&user, "ID = ?", "4d9e7fdc-a1ff-49db-92da-dc44e7395f35")
	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}

	user.Name = "Agnes"
	user.Email = "agnesdo@example.com"

	// Save the changes back to the database
	result = db.Save(&user)
	if result.Error != nil {
		panic("failed to update user: " + result.Error.Error())
	}

	log.Println("Пользователь — успешно обновлен:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.Name)
	log.Printf("	E-mail: <%s>\n", user.Email)
}

func DeleteUserByIdFromDB() {

	db := database.DB
	var user models.User
	user_id := "ec3236c5-4543-4fbe-bce0-ece27dded172"

	result := db.First(&user, "ID = ?", user_id)
	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}

	result = db.Delete(&user)
	if result.Error != nil {
		panic("failed to delete user: " + result.Error.Error())
	} else if result.RowsAffected == 0 {
		panic("no user record was deleted")
	} else {
		log.Println("Пользователь — успешно удален:")
		log.Printf("	ID пользователя: <%s>\n", user_id)

	}
}
