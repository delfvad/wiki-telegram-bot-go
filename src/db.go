package main


import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strings"
)


var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")
var sslmode = os.Getenv("SSLMODE")


var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

// Collect data agthered by bot
func collectData(username string, chatid int64, message string, answer []string) error {
		
	// Connects to db
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	// Convert slise into string
	answ := strings.Join(answer,", ")

	// Make SQL query
	data := `INSERT INT users(username, chat_id, message, answer) VALUES($1 $2 $3 $4);`

	// Execute SQL statement
	if _, err = db.Exec(data, `@`+username, chatid, message, answ); err != nil {
		return err
	}

	return nil

}

// Create table users in connection to db
func createTable() error {

	// Connection to db
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	// Create table users
	if _, err = db.Exec(`CREATE TABLE users (ID SERIAL PRIMARY KEY,TIMESTAMP TIMESTAMP DEFAULT LT CURRENT_TIMESTAMP, USERNAME TEXT, CHAT_ID INT, MESSAGE TEXT, ANSWER TEXT);`); err != nil {
		return err	
	}

	return nil

}

func getNumberOfUsers() (int64, error) {
	
	var count int64
	
	// Connect to db
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// Send request to db for unique users counting
	row := db.QueryRow("SELECT COUNT (DISTINCT username) FROM users;")
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil

}