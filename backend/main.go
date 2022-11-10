package main

import (
	// "fmt"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type EmailRequestBody struct {
	Mail string `json:"mail"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type User struct {
	Id   int64
	Mail string
}

// ! Esto es una cutrada para poder utilizar el contexto de conexión de la BBDD desde cualquier parte
var canyes *sql.DB

func main() {
	// Starts db connection
	dbConnection()

	// Initialises a router with the default functions.
	router := gin.Default()

	router.POST("/login", func(context *gin.Context) {
		// Set the struct requestBody must follow
		var requestBody EmailRequestBody

		if err := context.BindJSON(&requestBody); err != nil {
			// DO SOMETHING WITH THE ERROR
			context.String(http.StatusOK, "ERROR: Input not valid")
		}

		loginStatus := login(requestBody.Mail)

		context.JSON(http.StatusOK, loginStatus)
	})

	router.POST("/logout", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})

	// starts the server at port 8080
	router.Run(":8080")

}

// * API endpoints handlers

func login(mail string) LoginResponse {
	// ? creates and return valid session token

	// Queries db, rows is returned as pointers
	rows, err := canyes.Query(`SELECT * FROM "users" WHERE "mail" = $1`, mail)
	if err != nil {
		log.Fatal(err)
		return LoginResponse{false, ""}
	}
	defer rows.Close()

	// Initializes users array
	var users []User

	// Loop through rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.Id, &usr.Mail); err != nil {
			fmt.Println(err)
		}
		users = append(users, usr)
	}
	fmt.Println(users)
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}

	result := LoginResponse{true, "rge"}
	return result
}

func logout(mail string) {
	// Removes session token from table
}

func newInteraction(mail string, emoji string) {
	// Add new interaction
}

func checkSession(mail string, token string) {
	// if mail is logged then return true
	// else false
}

// * DB functions

func dbConnection() {
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Fatal(err)
	}

	// if db isnt working exit statuscode 2
	if err = db.Ping(); err != nil {
		panic(err)
	}
	canyes = db
	fmt.Println("The database is connected")
}

// * Misc functions

func generateToken() {

}
