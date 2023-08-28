package infrastructure

import (
	"PilotSoul/dynamic_service/src/domain"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Db *gorm.DB
}

var DB SqlHandler

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &domain.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DBName"),
	}

	db, err := domain.NewConnection(config)

	log.Println("running migrations")
	db.AutoMigrate(&domain.User{}, &domain.Segment{}, &domain.UserSegment{})
	db.SetupJoinTable(&domain.User{}, "Segments", &domain.UserSegment{})

	if err != nil {
		panic(err.Error)
	}

	// sqlHandler := new(SqlHandler)
	// sqlHandler.Db = db

	DB = SqlHandler{
		Db: db,
	}
}
