package conf

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

func ConnectDB() (DB *sql.DB, err error) {
	host := os.Getenv("POSTGRES_HOST")
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	username := os.Getenv("POSTGRES_USERNAME")
	dbName := os.Getenv("POSTGRES_DBNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	sslmode := os.Getenv("POSTGRES_SSLMODE")
	var connstring = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, dbName, sslmode)

	//connstring := "user=postgres dbname=postgres password='root' host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		panic(err)
	}
	return db, err
}
