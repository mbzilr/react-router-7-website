package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import as a side effect for PostgreSQL driver
)

type User struct {
	Name  string `db:"username"`
	Email string `db:"email"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbSSLMode)

	// Connect to a PostgreSQL Database using env.
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Database connection error:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	} else {
		log.Println("Successfully connected!")
	}

	var users []User
	rows, err := db.Query("SELECT username, email FROM users")
	if err != nil {
		log.Fatalln("Query error:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Name, &user.Email); err != nil {
			log.Fatalln("Row scan error:", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatalln("Rows iteration error:", err)
	}

	log.Printf("Retrieved users: %+v\n", users)
}
