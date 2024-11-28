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
	log.Println("Successfully connected to the database")
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
		log.Printf("Error fetching Toronto time: %v", err)
		return
	}

	// Log the current time in the database
	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", torontoTime)
	if err != nil {
		http.Error(w, "Error logging time", http.StatusInternalServerError)
		log.Printf("Error logging time: %v", err)
		return
	}

	// Send the current time as a JSON response
	response := map[string]string{"current_time": torontoTime.Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	log.Printf("Returned current time: %v", torontoTime)
}

// Handler to fetch all logged timestamps
func timeLogsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, timestamp FROM time_log")
	if err != nil {
		http.Error(w, "Error retrieving logs", http.StatusInternalServerError)
		log.Printf("Error retrieving time logs: %v", err)
		return
	}
	defer rows.Close() // Ensure rows are closed after processing

	var logs []map[string]interface{}
	for rows.Next() {
		var id int
		var timestamp string
		if err := rows.Scan(&id, &timestamp); err != nil {
			log.Printf("Error scanning row: %v", err)
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

	log.Printf("Returned %d time logs", len(logs))
}

// New handler to fetch all logged times in JSON format
func allLoggedTimesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, timestamp FROM time_log")
	if err != nil {
		http.Error(w, "Error retrieving all logs", http.StatusInternalServerError)
		log.Printf("Error retrieving all time logs: %v", err)
		return
	}
	defer rows.Close() // Ensure rows are closed after processing

	var logs []map[string]interface{}
	for rows.Next() {
		var id int
		var timestamp string
		if err := rows.Scan(&id, &timestamp); err != nil {
			log.Printf("Error scanning row: %v", err)
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

	log.Printf("Returned all time logs: %d records", len(logs))
}

// Set up logging to a file and stdout
func setupLogging() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to set up logging: %v", err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile)) // Write logs to both file and console
	log.Println("Logging initialized")
}

// Log incoming HTTP requests
func logRequest(r *http.Request) {
	log.Printf("Request: %s %s", r.Method, r.URL.Path)
}

func main() {
	setupLogging()
	initDB()
	defer db.Close() // Close database connection when the application exits

	http.HandleFunc("/current-time", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r) // Log request details
		currentTimeHandler(w, r)
	})

	http.HandleFunc("/time-logs", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r) // Log request details
		timeLogsHandler(w, r)
	})

	// New endpoint to get all logged times
	http.HandleFunc("/all-logged-times", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r) // Log request details
		allLoggedTimesHandler(w, r)
	})

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the HTTP server
}
