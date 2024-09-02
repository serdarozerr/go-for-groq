package main

import (
	"context"
	"github.com/serdarozerr/go-for-groq/internal/driver"
	"github.com/serdarozerr/go-for-groq/internal/service"
	"log"
	"net/http"
)

func main() {

	mongo_client := driver.GetClient()
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer driver.CloseClient(mongo_client)

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		service.CreateUser(w, r, mongo_client)
	})

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		service.MakeGroqRequest(w, r, mongo_client)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}

}
