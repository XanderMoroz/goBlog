package controllers

import (
	// "fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
	"github.com/XanderMoroz/goBlog/internal/utils"
)

// func Some() {

// 	pass := "мой пароль"

// 	// Хэшируем пароль
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Println("Ghjjbkdhvh")
// 	}

// 	log.Println("мой пароль:", pass)
// 	log.Println("Хэш пароля:", hashedPassword)

// 	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass))
// 	if err != nil {
// 		log.Println("Invalid Password:", err)
// 	}
// 	log.Println("... успешно")
// }

// @Summary        user registration
// @Description    Register User in app with given request body
// @Tags           Authentication
// @Accept         json
// @Produce        json
// @Param          request         	body        models.SignUpUserRequest    true    "Введите данные для регистрации"
// @Success        201              {string}    map[string]string
// @Failure        400              {string}    string    "Bad Request"
// @Router         /api/v1/register 			[post]
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
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
	log.Printf("	hashedPassword: <%+v>\n", hashedPassword)

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
// @Tags           Authentication
// @Accept         json
// @Produce        json
// @Param          request         	body        models.LoginRequest    true    "Введите данные для авторизации"
// @Success        201              {string}    map[]
// @Failure        400              {string}    string    "Bad Request"
// @Router         /api/v1/login 			[post]
func Login(c *fiber.Ctx) error {
	log.Println("Получен запрос на аутентификацию пользователя")

	db := database.DB

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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(body.Password))
	if err != nil {
		log.Println("Invalid Password:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}
	log.Println("... успешно")

	log.Println("Генерируем токен доступа...")
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Println("Error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	} else {
		log.Println("... успешно:", token)
	}

	log.Println("Устанавливаем токен в куки пользователя...")
	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24), // Expires in 24 hours
		// HTTPOnly: true,
		// Secure:   true,
	}
	c.Cookie(&cookie)
	log.Println("... успешно")

	// Аутентификация прошла успешно, возврящаем ответ
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":            "Login successful",
		"user_authenticated": user,
	})
}

// @Summary		get current user
// @Description Get token from users cookee
// @Tags 		Authentication
// @ID			get-current-user
// @Produce		json
// @Success		200		{object}	[]models.UserResponse
// @Router		/api/v1/current_user [get]
func GetCurrentUser(c *fiber.Ctx) error {
	log.Println("Получен запрос на извлечение авторизованного пользователя")

	// db := database.DB

	// Извлекаем JWT токен из куки пользователя
	cookieWithJWT := c.Cookies("jwt")

	log.Println("Извлекаем ID пользователя по из JWT токена")
	userID, err := utils.ParseUserIDFromJWTToken(cookieWithJWT)

	if err != nil {
		log.Println("Ошибка:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "Please Log in",
		})
	} else {
		log.Println("USER_ID из токена:", userID)
	}

	log.Println("Извлекаем пользователя по ID...")
	user := utils.GetUserByIDFromDB(userID)

	// Возвращаем данные пользователя
	return c.JSON(fiber.Map{
		"message":      "User retrieved successful",
		"current_user": user,
	})
}

// @Summary		logout current user
// @Description Clear JWT token by setting an empty value and expired time in the cookie
// @Tags 		Authentication
// @ID			logout-current-user
// @Produce		json
// @Success		200		{string}	map[]
// @Router		/api/v1/logout [get]
func Logout(c *fiber.Ctx) error {
	log.Println("Получен запрос на выход из аккаунта")

	log.Println("Удаляем токен из куков...")
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Expired 1 hour ago
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)
	log.Println("... успешно")

	// Return success response indicating logout was successful
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Logout successful",
	})
}
