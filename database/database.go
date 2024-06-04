package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/XanderMoroz/goBlog/internal/models"
)

type AppEnvConfig struct {
	AppSecret  string
	Dbdriver   string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

// Указатель на БД
// (он будет осуществлять запросы)
var DB *gorm.DB

// Извлкает переменные окружения и складывает в DBEnvConfig
func GetEnvConfig() *AppEnvConfig {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Не удалось извлечь .env")
	}

	return &AppEnvConfig{
		AppSecret: os.Getenv("API_SECRET"),

		Dbdriver:   os.Getenv("DB_DRIVER"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
	}
}

// Подключается к БД
func Connect() {

	envConfig := GetEnvConfig()

	DBURL := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		envConfig.Dbdriver,
		envConfig.DbUser,
		envConfig.DbPassword,
		envConfig.DbHost,
		envConfig.DbPort,
		envConfig.DbName)

	log.Printf("Подключаемся к БД: <%s> ...", envConfig.Dbdriver)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		log.Println("Невозможно подключиться к базе ")
		log.Fatal("connection error:", err)
	} else {
		log.Println("... успешно")
		log.Printf("	DB_HOST: <%s>", envConfig.DbHost)
		log.Printf("	DB_PORT: <%s>", envConfig.DbPort)
		log.Printf("	DB_NAME: <%s>", envConfig.DbName)
		log.Printf("	DB_USER: <%s>", envConfig.DbUser)
	}

	DB = db

	log.Printf("Устанавливаем миграции в БД...")
	db.AutoMigrate(&models.User{}, &models.Article{}, &models.Category{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
	log.Printf("... успешно!")
}
