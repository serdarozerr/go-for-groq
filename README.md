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


