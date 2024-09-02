# GO for Groq

This backend is designed to interact with Groq-hosted LLM models, following a query-response logic. It utilizes Groq's
high-speed inference capabilities and Go's goroutines to deliver rapid responses and efficiently scale to handle
numerous users. The purpose of this backend is to demonstrate and utilize the Groq client within the Go programming
language. Currently, Groq does not offer an official Go library.

## Specs

- Database: MongoDB is integrated, with the Go MongoDB client being utilized. This client is thread-safe, allowing for
  seamless use across goroutines in Go.
- LLM: The backend leverages Groq-hosted Mistral models for large language model (LLM) tasks.
- Backend: The backend is built using Go, taking advantage of its concurrency features and performance.

## Pre Requisites

You need to env variable should to exported in your environment:

- GROQ_API_KEY
- MONGO_URI

Export these keys in your terminal like following:

      export GROQ_API_KEY="<your api key>"
      export MONGO_URI="<your uri>"


## How to Use
Once the building is done run the backend:

    go build cmd/go-for-groq/main.go (build)
    ./main (run)

- First, Create a user and save the returned user_id to ask query

       curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{"name": "John", "surname":Doe"}'  

- Second, Send query

       curl -X POST http://localhost:8080/query -H "Content-Type: application/json" -d '{"query": "who is the founder of Manner waffer?" , "user_id":"<your user id>"}'


