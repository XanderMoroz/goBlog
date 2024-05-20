package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
)

// @Summary        create new user
// @Description    Creating User in DB with given request body
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          request            body        models.CreateUserRequest    true    "Request of Creating User Object"
// @Success        201                {string}    string
// @Failure        400                {string}    string    "Bad Request"
// @Router         /users [post]
func AddNewUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)

	// Store the body in the note and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	log.Println("Запрос успешно обработан обработан")

	// Add a uuid to the note
	user.ID = uuid.New()
	// Create the Note and return error if encountered
	log.Printf("Добавляем в БД нового пользователя:")
	log.Printf("	ID: <%+v>\n", user.ID)
	log.Printf("	Имя: <%+v>\n", user.FirstName)

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	log.Println("Новый пользователь — успешно создан:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.FirstName)

	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "User Created", "data": user})
}

// @Summary		get all users
// @Description Get all users from db
// @Tags 		Users
// @ID			get-all-users
// @Produce		json
// @Success		200		{object}	[]models.UserResponse
// @Router		/users [get]
func GetAllUsers(c *fiber.Ctx) error {

	db := database.DB
	var users []models.User // users slice

	result := db.Find(&users)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve users: " + result.Error.Error())
	}

	log.Println("Список пользователей — успешно извлечен:")
	for _, user := range users {
		log.Printf("User ID: %s, Name: %s %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
	}

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No notes present",
			"data":    nil,
		})
	}

	c.Status(http.StatusOK)
	// return c.JSON(articleList)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Notes Found",
		"data":    users,
	})
}

// @Summary		get a user item by ID
// @Description Get a user item by ID
// @Tags 		Users
// @ID			get-user-by-id
// @Produce		json
// @Param		id	path		string	true	"userUUID"
// @Success		200	{object}	models.UserResponse
// @Failure		404	{object}	[]string
// @Router		/users/{id} [get]
func GetUserById(c *fiber.Ctx) error {

	db := database.DB
	var user models.User

	// Read the param userUUID
	id := c.Params("id")

	// Retrieve the record you want to update
	result := db.First(&user, "ID = ?", id)

	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}

	// If no such note present return an error
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	log.Println("Пользователей — успешно извлечен:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.FirstName)
	log.Printf("	Фамилия: <%s>\n", user.LastName)
	log.Printf("	E-mail: <%s>\n", user.Email)

	// Return the note with the Id
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User Found",
		"data":    user,
	})
}

// @Summary		delete a user item by ID
// @Description Delete a user item by ID
// @ID			delete-user-by-id
// @Tags 		Users
// @Produce		json
// @Param		id	path		string	true	"userUUID"
// @Success		200	{object}	[]string
// @Failure		404	{object}	[]string
// @Router		/users/{id} [delete]
func DeleteUserById(c *fiber.Ctx) error {

	db := database.DB
	var user models.User

	// Read the param userUUID
	id := c.Params("id")

	// Retrieve the record you want to dekete
	result := db.First(&user, "ID = ?", id)

	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}

	// If no such note present return an error
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User Not Found", "data": nil})
	}

	// Delete the note and return error if encountered
	err := db.Delete(&user, "ID = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	log.Println("Пользователь — успешно удален:")
	log.Printf("	ID: <%s>\n", user.ID)

	// Return success message
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User Deleted",
	})
}

func CreateUserInDB() {

	db := database.DB
	newUser := models.User{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane111doe@gmail.com",
		Country:   "Spain",
		Role:      "Chef",
		Age:       30,
	}

	// Add a uuid to a new user record...
	newUser.ID = uuid.New()
	// ... Create a new user record...
	result := db.Create(&newUser)
	if result.Error != nil {
		panic("failed to create user: " + result.Error.Error())
	}
	// ... Handle successful creation ...
	log.Println("Новый пользователь — успешно создан:")
	log.Printf("	ID: <%s>\n", newUser.ID)
	log.Printf("	Имя: <%s>\n", newUser.FirstName)
}

func GetUsersFromDB() {

	db := database.DB
	var users []models.User // users slice

	result := db.Find(&users)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve users: " + result.Error.Error())
	}

	log.Println("Список пользователей — успешно извлечен:")
	for _, user := range users {
		log.Printf("User ID: %s, Name: %s %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
	}

}

func GetUserByIdFromDB() {

	db := database.DB
	var users []models.User // users slice

	result := db.Where("ID = ?", "073b515e-f9b6-45b2-aeb1-bbaa6de48286").Find(&users)

	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}
	// iterate over the users slice and print the details of each user
	log.Println("Пользователь — успешно извлечен:")
	for _, user := range users {
		log.Printf("	ID: <%s>\n", user.ID)
		log.Printf("	Имя: <%s>\n", user.FirstName)
		log.Printf("	Фамилия: <%s>\n", user.LastName)
		log.Printf("	E-mail: <%s>\n", user.Email)
	}

}

func GetUserByNameFromDB() {

	db := database.DB
	var users []models.User // users slice

	result := db.Where("First_Name = ?", "Jane").Find(&users)

	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}
	// iterate over the users slice and print the details of each user
	log.Println("Пользователь — успешно извлечен:")
	for _, user := range users {
		log.Printf("	ID: <%s>\n", user.ID)
		log.Printf("	Имя: <%s>\n", user.FirstName)
		log.Printf("	Фамилия: <%s>\n", user.LastName)
		log.Printf("	E-mail: <%s>\n", user.Email)
	}

}

func UpdateUserByIdFromDB() {

	db := database.DB

	var user models.User

	// Retrieve the record you want to update
	result := db.First(&user, "ID = ?", "ec3236c5-4543-4fbe-bce0-ece27dded172")
	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}

	user.FirstName = "Agnes"
	user.LastName = "Doe"
	user.Email = "agnesdo@example.com"

	// Save the changes back to the database
	result = db.Save(&user)
	if result.Error != nil {
		panic("failed to update user: " + result.Error.Error())
	}

	log.Println("Пользователь — успешно обновлен:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.FirstName)
	log.Printf("	Фамилия: <%s>\n", user.LastName)
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
