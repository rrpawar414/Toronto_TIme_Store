package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize the database connection
func initDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	db, err = sql.Open("mysql", dsn)
	if err != nil || db.Ping() != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
}

// Get the current time in Toronto's timezone
func getTorontoTime() (time.Time, error) {
	location, err := time.LoadLocation("America/Toronto") // Load Toronto timezone
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(location), nil
}

// Handler to return current Toronto time in JSON format
func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	torontoTime, err := getTorontoTime()
	if err != nil {
		http.Error(w, "Error fetching time", http.StatusInternalServerError)
		return
	}

	// Log the current time in the database
	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", torontoTime)
	if err != nil {
		http.Error(w, "Error logging time", http.StatusInternalServerError)
		return
	}

	// Send the current time as a JSON response
	response := map[string]string{"current_time": torontoTime.Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler to fetch all logged timestamps
func timeLogsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, timestamp FROM time_log")
	if err != nil {
		http.Error(w, "Error retrieving logs", http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Ensure rows are closed after processing

	var logs []map[string]interface{}
	for rows.Next() {
		var id int
		var timestamp string
		if err := rows.Scan(&id, &timestamp); err != nil {
			continue // Skip rows with errors
		}

		// Convert timestamp to RFC3339 format
		parsedTime, _ := time.Parse("2006-01-02 15:04:05", timestamp)
		logs = append(logs, map[string]interface{}{
			"id":        id,
			"timestamp": parsedTime.Format(time.RFC3339),
		})
	}

	// Send the logs as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

// Set up logging to a file and stdout
func setupLogging() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to set up logging: %v", err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile)) // Write logs to both file and console
}

func main() {
	setupLogging()
	initDB()
	defer db.Close() // Close database connection when the application exits

	http.HandleFunc("/current-time", currentTimeHandler) // Endpoint for current time
	http.HandleFunc("/time-logs", timeLogsHandler)       // Endpoint for time logs

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the HTTP server
}
