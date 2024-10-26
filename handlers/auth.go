package handlers

import (
	"AUTHAPI/models"
	"AUTHAPI/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := models.userCollection.InsertOne(ctx, user)

	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	err := models.userCollection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(user.Email)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
