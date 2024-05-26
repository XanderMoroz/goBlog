package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
)

// @Summary        user registration
// @Description    Register User in DB with given request body
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          request         	body        models.SignUpUserRequest    true    "Введите данные для регистрации"
// @Success        201              {string}    string
// @Failure        400              {string}    string    "Bad Request"
// @Router         /register 			[post]
func Register(c *fiber.Ctx) error {

	log.Println("Получен запрос на регистрацию пользователя")

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

	var existingUser models.User
	// Проверяем уникальность данных email
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		log.Println("Этот email уже занят. Попробуйти ввести другой")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Этот email уже занят. Попробуйти ввести другой",
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Email), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Присваиваем новому пользователю уникальный ID
	user.ID = uuid.New()
	user.Password = hashedPassword

	log.Printf("Добавляем в БД нового пользователя:")
	log.Printf("	ID: <%+v>\n", user.ID)
	log.Printf("	Имя: <%+v>\n", user.Name)
	log.Printf("	Email: <%+v>\n", user.Email)

	// Создаем пользователя в БД
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	log.Println("Новый пользователь — успешно создан:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.Name)

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})
}
