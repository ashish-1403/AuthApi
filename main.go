package main

import (
	"AUTHAPI/config"
	"AUTHAPI/handlers"
	"AUTHAPI/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize MongoDB connection
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	models.SetDatabase(db)

	// Router setup
	r := mux.NewRouter()
	r.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/signin", handlers.SignIn).Methods("POST")
	// r.HandleFunc("/refresh", handlers.RefreshToken).Methods("POST")
	// r.HandleFunc("/revoke", handlers.RevokeToken).Methods("POST")

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
