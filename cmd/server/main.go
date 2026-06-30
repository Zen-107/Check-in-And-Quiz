package main

import (
	"check-in-and-quiz/internal/database"
	"check-in-and-quiz/internal/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Get database configuration from environment
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "checkin_quiz"
	}

	// Initialize database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	err := database.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to database")

	// ไม่ต้อง migrate เพราะใช้ init.sql ใน Docker แล้ว
	// Tables are created by init.sql in Docker volume

	// Setup router
	r := mux.NewRouter()

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/checkin", handlers.CheckInHandler).Methods("GET", "POST")
	r.HandleFunc("/quiz", handlers.QuizHandler).Methods("GET", "POST")
	r.HandleFunc("/result", handlers.ResultHandler).Methods("GET")
	r.HandleFunc("/api/checkin/count", handlers.GetCheckInCountHandler).Methods("GET")
	r.HandleFunc("/api/checkin", handlers.SubmitCheckInHandler).Methods("POST")
	r.HandleFunc("/api/quiz", handlers.SubmitQuizHandler).Methods("POST")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
