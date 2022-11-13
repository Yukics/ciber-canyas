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
	"github.com/go-co-op/gocron"
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

type InteractionRequestBody struct {
	Mail  string `json:"mail"`
	Token string `json:"token"`
	Emoji string `json:"emoji"`
}

// * HTTP Responses

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type LogoutResponse struct {
	Success bool `json:"success"`
}

type InteractionResponse struct {
	Success bool `json:"success"`
}

// * SQL structs

type User struct {
	Id   int
	Mail string
}

type Session struct {
	Id         int
	Mail       string
	Expiration time.Time
	Token      string
}

type Emoji struct {
	Emoji string `json:"emoji"`
	Count int64  `json:"count"`
}

type Interactor struct {
	Mail  string `json:"mail"`
	Count int    `json:"count"`
}

// ! Esto es una cutrada para poder utilizar el contexto de conexión de la BBDD desde cualquier parte
var canyes *sql.DB

func main() {

	// Starts db connection
	dbConnection()

	s := gocron.NewScheduler(time.UTC)
	// Every 5 min
	s.Cron("*/5 * * * *").Do(func() {
		cleanSessions()
	})
	// starts the scheduler asynchronously
	s.StartAsync()

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

	router.GET("/emojis", func(context *gin.Context) {

		emojis := getEmojis()

		context.JSON(http.StatusOK, emojis)
	})

	router.POST("/interaction", func(context *gin.Context) {
		// Set the struct requestBody must follow
		var requestBody InteractionRequestBody

		if err := context.BindJSON(&requestBody); err != nil {
			context.String(http.StatusOK, "ERROR: Input not valid")
		}

		interactionStatus := interact(requestBody)

		context.JSON(http.StatusOK, interactionStatus)
	})

	router.GET("/topInteractors", func(context *gin.Context) {

		interactions := getTopInteractors()

		context.JSON(http.StatusOK, interactions)
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

func logout(reqBody LogoutRequestBody) LogoutResponse { // * DONE
	// Removes session token from table
	sessionId := checkSession(reqBody.Mail, reqBody.Token)
	if sessionId == 0 {
		return LogoutResponse{false}
	}

	_, err := canyes.Exec(`DElETE FROM sessions WHERE id like $1)`, sessionId)
	if err != nil {
		fmt.Println(err)
		return LogoutResponse{false}
	}

	return LogoutResponse{true}
}

func interact(reqBody InteractionRequestBody) InteractionResponse {
	// Check if user is logged in
	sessionId := checkSession(reqBody.Mail, reqBody.Token)

	if sessionId == 0 {
		return InteractionResponse{false}
	}

	// Add new interaction
	_, err := canyes.Exec(`INSERT INTO interactions(user_id, element_id) (
		SELECT x.user_id, y.element_id 
		FROM 
			(select 'a' as joinner, u.user_id from users u where u.mail like $1) x
		left join 
			(select 'a' as joinner, e.element_id from elements e where e.emoji like $2) y
		on (x.joinner = y.joinner)
	)`, reqBody.Mail, reqBody.Emoji)

	if err != nil {
		fmt.Println(err)
		return InteractionResponse{false}
	}

	return InteractionResponse{true}
}

func getEmojis() []Emoji {
	// * Emoji codes: https://unicode.org/emoji/charts/full-emoji-list.html
	rows, err := canyes.Query(`SELECT e.emoji, COALESCE(COUNT(i.element_id), 0) as count FROM elements e LEFT JOIN interactions i ON i.element_id = e.element_id GROUP BY e.emoji;`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var emojis []Emoji

	for rows.Next() {
		var emj Emoji
		if err := rows.Scan(&emj.Emoji, &emj.Count); err != nil {
			fmt.Println(err)
			return []Emoji{{"", 0}}
		}
		emojis = append(emojis, emj)
	}

	return emojis
}

func getTopInteractors() []Interactor {
	rows, err := canyes.Query(`SELECT u.mail , COALESCE(COUNT(i.user_id), 0) as count FROM users u LEFT JOIN interactions i ON i.user_id  = u.user_id GROUP BY u.mail order by count;`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var interactors []Interactor

	for rows.Next() {
		var intr Interactor
		if err := rows.Scan(&intr.Mail, &intr.Count); err != nil {
			fmt.Println(err)
			return []Interactor{{"", 0}}
		}
		interactors = append(interactors, intr)
	}

	return interactors
}

func checkSession(mail string, token string) int {
	// if mail is logged then return id
	// else 0

	rows, err := canyes.Query(`SELECT session_id,expiration,mail FROM sessions s INNER JOIN users u ON u.user_id = s.user_id WHERE u.mail like $1 AND s.token like $2 AND  s.expiration > now() AT time ZONE 'utc';`, mail, token)

	if err != nil {
		log.Fatal(err)
		return 0
	}

	defer rows.Close()

	// Initializes sessions array
	var sessions []Session

	// Serialize data
	for rows.Next() {
		var sesion Session
		if err := rows.Scan(&sesion.Id, &sesion.Expiration, &sesion.Mail); err != nil {
			fmt.Println(err)
			return 0
		}
		sessions = append(sessions, sesion)
	}

	if len(sessions) >= 1 {
		return sessions[0].Id
	}

	return 0
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

func cleanSessions() {
	_, err := canyes.Exec(`DELETE FROM sessions WHERE expiration < now() AT time ZONE 'utc';`)
	if err != nil {
		fmt.Println(err)
	}
}

// * Misc functions

func generateToken() string {
	b := make([]byte, 64)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func generateExpiration() time.Time {
	// Small hack for not thinking about utc shit
	now := 60

	expiration := time.Now().Local().Add(time.Minute * time.Duration(now+60))
	return expiration
}
