package api

import _ "github.com/lib/pq"

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Simnek/web-go/middleware"
	"github.com/Simnek/web-go/types"
	"log"
	"net/http"
)

type postUserHandler struct{}

func (p postUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	HandlePostUser(w, r)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {

	var user types.User
	_ = user

	fmt.Fprintf(w, "user")
}

func HandlePostUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling POST request for /user/create")
	log.Println(r.Method)
	if r.Method == http.MethodPost {
		// Parse the request body
		var newUser types.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(newUser.Name)
		log.Println(newUser.Email)

		// Establish a connection to the PostgreSQL database
		db, err := sql.Open("postgres", "postgresql://postgres:postgres@10.21.59.29:5432/postgres?sslmode=disable&search_path=tksmvgo")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		sqlStatement := `INSERT INTO users (name, email) VALUES ($1, $2)`

		// Execute an SQL INSERT query
		_, err = db.Exec(sqlStatement, newUser.Name, newUser.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		fmt.Fprintf(w, "New user created successfully")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlePostUserWithCORS(w http.ResponseWriter, r *http.Request) {
	// Use the CORS middleware before calling HandlePostUser
	middleware.CORSHandler(postUserHandler{}).ServeHTTP(w, r)
}
