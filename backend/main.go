package main

import (
	// "fmt"
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// * HTTP Requests
type EmailRequestBody struct {
	Mail string `json:"mail"`
}

type LogoutRequestBody struct {
	Mail  string `json:"mail"`
	Token string `json:"token"`
}

// * HTTP Responses

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type LogoutResponse struct {
	Success bool `json:"success"`
}

// * SQL structs

type User struct {
	Id   int64
	Mail string
}

type Session struct {
	Id         int64
	Mail       string
	Expiration time.Time
	Token      string
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
			context.String(http.StatusOK, "ERROR: Input not valid")
		}

		loginStatus := login(requestBody.Mail)

		context.JSON(http.StatusOK, loginStatus)
	})

	router.POST("/logout", func(context *gin.Context) {
		// Set the struct requestBody must follow
		var requestBody LogoutRequestBody

		if err := context.BindJSON(&requestBody); err != nil {
			context.String(http.StatusOK, "ERROR: Input not valid")
		}

		logoutStatus := logout(requestBody)

		context.JSON(http.StatusOK, logoutStatus)
	})

	// starts the server at port 8080
	router.Run(":8080")

}

// * API endpoints handlers

func login(mail string) LoginResponse { // * DONE
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
			return LoginResponse{false, ""}
		}
		users = append(users, usr)
	}

	token := generateToken()
	expirate := generateExpiration()

	if len(users) == 1 {
		_, err := canyes.Exec(`INSERT INTO sessions (session_id,user_id,expiration,token) VALUES (DEFAULT, $1,$2,$3)`, users[0].Id, expirate, token)
		if err != nil {
			fmt.Println(err)
			return LoginResponse{false, ""}
		}
	}

	return LoginResponse{true, token}
}

func logout(body LogoutRequestBody) LogoutResponse { // * DONE
	// Removes session token from table

	rows, err := canyes.Query(`select session_id,expiration,token,mail from sessions s inner join users u ON u.user_id = s.user_id where u.mail like $1 AND s.token like $2;`, body.Mail, body.Token)
	if err != nil {
		log.Fatal(err)
		return LogoutResponse{false}
	}

	defer rows.Close()

	// Initializes users array
	var sessions []Session

	// Loop through rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var sesion Session
		if err := rows.Scan(&sesion.Id, &sesion.Expiration, &sesion.Mail); err != nil {
			fmt.Println(err)
			return LogoutResponse{false}
		}
		sessions = append(sessions, sesion)
	}
	if len(sessions) >= 1 {
		_, err := canyes.Exec(`DElETE FROM sessions WHERE id like $1)`, sessions[0].Id)
		if err != nil {
			fmt.Println(err)
			return LogoutResponse{false}
		}
	}

	return LogoutResponse{true}
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

func generateToken() string {
	b := make([]byte, 64)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func generateExpiration() time.Time {
	expiration := time.Now().Local().Add(time.Minute * time.Duration(60))
	return expiration
}
