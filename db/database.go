package database

import (
	"fmt"
	"log"
	"time"

	"pennywise-api/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx" // Using SQLX as an example, you can replace with your preferred DB driver
)

var DB *sqlx.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	// Build connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
	)

	// Connect to the database
	var err error
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Configure the connection pool
	DB.SetMaxOpenConns(25)                 // Maximum number of open connections to the database
	DB.SetMaxIdleConns(25)                 // Maximum number of idle connections in the pool
	DB.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection

	log.Println("Database connection successful!")
}
