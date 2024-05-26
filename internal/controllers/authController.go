package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
)

// @Summary        user registration
// @Description    Register User in app with given request body
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          request         	body        models.SignUpUserRequest    true    "Введите данные для регистрации"
// @Success        201              {string}    map[string]string
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

	// Хэшируем пароль
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
	log.Printf("	Email: <%+v>\n", user.Email)

	return c.JSON(fiber.Map{"status": "success", "message": "User registered successfully", "data": user})
}

// @Summary        user authentication
// @Description    Authenticate User in app with given request body
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          request         	body        models.LoginRequest    true    "Введите данные для авторизации"
// @Success        201              {string}    string
// @Failure        400              {string}    string    "Bad Request"
// @Router         /login 			[post]
func Login(c *fiber.Ctx) error {
	log.Println("Получен запрос на аутентификацию пользователя")

	db := database.DB
	log.Println("Received a Login request")

	// Parse request body
	body := new(models.LoginRequest)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	// Check if user exists
	var user models.User
	var defaulUserID uuid.UUID
	db.Where("email = ?", body.Email).First(&user)
	if user.ID == defaulUserID {
		log.Println("User not found")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	log.Println("Верифицируем пароль...")
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		log.Println("Invalid Password:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}
	log.Println("... успешно")

	log.Println("Генерируем токен доступа...")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})
	token, err := claims.SignedString([]byte("SomeAppSecret"))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}
	log.Println("... успешно")

	log.Println("Устанавливаем токен в куки пользователя...")
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // Expires in 24 hours
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)
	log.Println("... успешно")

	// Аутентификация прошла успешно, возврящаем ответ
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":            "Login successful",
		"user_authenticated": user,
	})
}
