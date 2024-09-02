package service

import (
	"encoding/json"
	"github.com/serdarozerr/go-for-groq/internal/db"
	"github.com/serdarozerr/go-for-groq/internal/driver"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	user := db.User{QueryAnswers: []db.QueryAnswer{}}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	collection := driver.GetCollection(client, "User")

	result, err := user.Create(collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := db.Response{Id: *result}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
