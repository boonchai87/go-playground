package conf

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
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
	fmt.Print(connstring)
	//connstring := "user=postgres dbname=postgres password='root' host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		panic(err)
	}
	return db, err
}
func mysqlDB() (DB *sql.DB, err error) {
	// connect db
	fmt.Print(os.Getenv("MYSQL_HOST") + "," + os.Getenv("MYSQL_PASS"))
	var mysql_user = os.Getenv("MYSQL_USER")
	var mysql_password = os.Getenv("MYSQL_PASSWORD")
	// https://go.dev/doc/tutorial/database-access
	cfg := mysql.Config{
		User:   mysql_user,
		Passwd: mysql_password,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "mydb",
	}
	db, err := sql.Open("postgres", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db, err
}
