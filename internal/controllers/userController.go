package controllers

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
// @Param          request         	body        models.CreateUserRequest    true    "Введите данные пользователя"
// @Success        201              {string}    string
// @Failure        400              {string}    string    "Bad Request"
// @Router         /users 			[post]
func AddNewUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)

	// Извлекаем тело запроса
	err := c.BodyParser(user)
	if err != nil {
		// Обрабатываем ошибку
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Проверьте данные",
			"data":    err,
		})
	}
	log.Println("Запрос успешно обработан обработан")

	// Add a uuid to the note
	user.ID = uuid.New()

	log.Printf("Добавляем в БД нового пользователя:")
	log.Printf("	ID: <%+v>\n", user.ID)
	log.Printf("	Имя: <%+v>\n", user.Name)

	// Создаем пользователя
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	log.Println("Новый пользователь — успешно создан:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.Name)

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
		log.Printf("User ID: <%s>, Name: <%s>, Email: <%s>\n", user.ID, user.Name, user.Email)
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

// @Summary		get a user by ID
// @Description Get a user by ID
// @Tags 		Users
// @ID			get-user-by-id
// @Produce		json
// @Param		id	path		string	true	"userUUID"
// @Success		200	{object}	models.UserResponse
// @Failure		404	{object}	[]string
// @Router		/users/{id} 	[get]
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
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	log.Println("Пользователь — успешно извлечен:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.Name)
	log.Printf("	E-mail: <%s>\n", user.Email)

	// Return the note with the Id
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User Found",
		"data":    user,
	})
}

// @Summary			update user by ID
// @Description 	Update user by ID
// @ID				delete-user-by-id
// @Tags 			Users
// @Produce			json
// @Param			id		path		string	true	"userUUID"
// @Param          request         	body        models.UpdateUserBody    true    "Введите данные пользователя"
// @Success			200	{object}	[]string
// @Failure			404	{object}	[]string
// @Router			/users/{id} [put]
func UpdateUserById(c *fiber.Ctx) error {

	db := database.DB
	body := new(models.UpdateUserBody)
	var user models.User

	// Read the param userUUID
	id := c.Params("id")

	// Извлекаем запись по ID
	result := db.First(&user, "ID = ?", id)

	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}

	// If no such note present return an error
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	log.Println(id)

	// Извлекаем тело запроса
	err := c.BodyParser(body)
	if err != nil {
		// Обрабатываем ошибку
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Проверьте данные",
			"data":    err,
		})
	}
	log.Println("Тело запроса извлечено:")
	log.Printf("	Новое Имя: <%s>\n", body.Name)
	log.Printf("	Новый E-mail: <%s>\n", body.Email)

	user.Name = body.Name
	user.Email = body.Email

	// Save the changes back to the database
	result = db.Save(&user)
	if result.Error != nil {
		panic("failed to update user: " + result.Error.Error())
	}

	log.Println("Пользователь — успешно обновлен:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.Name)
	log.Printf("	E-mail: <%s>\n", user.Email)

	// Return success message
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User Updated",
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
