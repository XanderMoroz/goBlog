package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
)

func Some() {

	pass := "мой пароль"

	// Хэшируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Ghjjbkdhvh")
	}

	log.Println("мой пароль:", pass)
	log.Println("Хэш пароля:", hashedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass))
	if err != nil {
		log.Println("Invalid Password:", err)
	}
	log.Println("... успешно")
}

// @Summary        user registration
// @Description    Register User in app with given request body
// @Tags           Authentication
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
// @Router         /login 			[post]
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
	log.Println("user.ID:", user.ID)
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
	log.Println("... успешно:")
	log.Println("... token:", token)

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
// @Router		/current_user [get]
func GetCurrentUser(c *fiber.Ctx) error {
	log.Println("Получен запрос на извлечение авторизованного пользователя")

	db := database.DB

	// Извлекаем JWT токен из куки пользователя
	cookie := c.Cookies("jwt")

	// Parse JWT token with claims
	hmacSecretString := "SomeAppSecret"
	hmacSecret := []byte(hmacSecretString)

	log.Println("Извлекаем токен из куки пользователя...")
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})
	// Handle token parsing errors
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse token",
		})
	} else {
		log.Println("... успешно")
	}

	log.Println("Извлекаем USER_ID из токена...")
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("Invalid JWT Token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse claims",
		})
	} else {
		log.Println("... успешно")
	}

	var user models.User
	log.Println("Извлекаем пользователя по ID...")
	result := db.Where("ID =?", claims["sub"]).First(&user)

	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	} else {
		log.Println("Пользователь — успешно извлечен:")
		log.Printf("	ID: <%s>\n", user.ID)
		log.Printf("	Имя: <%s>\n", user.Name)
		log.Printf("	E-mail: <%s>\n", user.Email)
	}

	// Return user details as JSON response
	return c.JSON(fiber.Map{
		"message":      "User retrieved successful",
		"current_user": user,
	})
}
