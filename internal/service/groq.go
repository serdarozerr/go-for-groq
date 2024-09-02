package service

import (
	"context"
	"encoding/json"
	"github.com/serdarozerr/go-for-groq/internal/db"
	"github.com/serdarozerr/go-for-groq/internal/driver"
	"github.com/serdarozerr/go-for-groq/internal/groq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func updateUserQueryAnswer(client *mongo.Client, groqQuery groq.GroqUserQuery, groqAnswer groq.GroqUserResponse) error {

	queryAnswer := db.QueryAnswer{Query: groqQuery.Query, Answer: groqAnswer.Answer}
	collection := driver.GetCollection(client, "User")
	userId, _ := primitive.ObjectIDFromHex(groqQuery.UserId)
	filter := bson.M{"_id": userId}
	update := bson.M{
		"$push": bson.M{
			"query_answers": queryAnswer,
		},
	}
	count, err := collection.UpdateOne(context.Background(), filter, update)
	println(count.MatchedCount)
	if err != nil {
		return err
	}

	return nil

}

func MakeGroqRequest(w http.ResponseWriter, r *http.Request, client *mongo.Client) {

	message := groq.GroqAPIQuery{}
	groqQuery := groq.GroqUserQuery{}
	err := json.NewDecoder(r.Body).Decode(&groqQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	answer, err := message.Request(groqQuery.Query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	groqAnswer := groq.GroqUserResponse{Answer: answer}
	groqAnswerJson, err := json.Marshal(groqAnswer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = updateUserQueryAnswer(client, groqQuery, groqAnswer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(groqAnswerJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
