package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
func Connect() {

	//load .env file
	err := godotenv.Load("internal/config/.env")
	CheckError(err) // log.Fatal("Error loading .env:", err)

	//Read .env variable
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	port := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s ", host, port, user, password, dbname, sslmode)

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	// DB, err = sql.Open("postgres", connectionString)
	CheckError(err) // log.Fatal("Error Opening DB :", err)	
	fmt.Println("Succesfully Connected to PostgreSql...")

}
