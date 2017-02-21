package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"project/app"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbMode := os.Getenv("DB_MODE")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, dbMode)

	var err error
	app.DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("DB error: ", err.Error())
	}

	if err = app.DB.Ping(); err != nil {
		log.Fatal("DB error: ", err.Error())
	} else {
		log.Print("DB: Connected")
	}

	port := os.Getenv("APP_PORT")
	log.Printf("App started inside container. Listening on %s", port)

	if err := app.ListenAndServe(port); err != nil {
		log.Fatal(err.Error())
	}
}
