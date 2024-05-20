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

// указатель на бд (он будет осуществлять запросы)
var DB *gorm.DB

// Инициализация базы данных
func Connect() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Не удалось извлечь .env")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")

	DBURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s", Dbdriver, DbUser, DbPassword, DbHost, DbPort, DbName)

	log.Printf("Осуществляем подключение к БД: <%s>", Dbdriver)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		log.Println("Невозможно подключиться к базе ")
		log.Fatal("connection error:", err)
	} else {
		log.Println("Подключение к БД прошло успешно")
		log.Printf("База данных: <%s> Хост: <%s> Порт: <%s>", DbName, DbHost, DbPort)
	}

	DB = db

	log.Println("Устанавливаем миграции в БД...")
	db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
	log.Println("Миграции установлены успешно")
}

// func Init() *gorm.DB {
// 	err := godotenv.Load("./envs/.env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	Dbdriver := "POSTGRES"
// 	DbUser := os.Getenv("DB_USERNAME")
// 	DbPassword := os.Getenv("DB_PASSWORD")
// 	DbHost := os.Getenv("DB_HOST")
// 	DbPort := os.Getenv("DB_PORT")
// 	DbName := os.Getenv("DB_NAME")

// 	DBURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s", DbUser, DbUser, DbPassword, DbHost, DbPort, DbName)

// 	log.Println("Осуществляем подключение к базе")
// 	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println("Невозможно подключиться к базе ", Dbdriver)
// 		log.Fatal("connection error:", err)
// 	} else {
// 		fmt.Println("Подключение к базе прошло успешно ", Dbdriver)
// 	}

// 	db.AutoMigrate(&models.Dog{})

// 	return db
// }

//
