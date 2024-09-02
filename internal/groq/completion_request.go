package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	systemMessageContent = "You are a highly efficient assistant that helps generate answer for the given questions.\nAnswers should be well-crafted, directly relevant and include reasoning.\nConsider different aspects and angles related to the query to provide comprehensive coverage.\nBad Answers will be penalized.Ensure the JSON object follows this schema: {answer: answer}"
	userMessageContent   = "You are a highly efficient assistant. Please analyze the following user query and create answer for it, the search query is : %s \nDuring the creating answer understand the whole context and understand what user is exactly looking."
)

func (mes GroqAPIQuery) Request(query string) (string, error) {

	systemMessage := Message{Role: "system", Content: systemMessageContent}
	userMessage := Message{Role: "user", Content: fmt.Sprintf(userMessageContent, query)}

	grogAPIMessage := GroqAPIQuery{}
	messages := append(grogAPIMessage.Messages, systemMessage, userMessage)
	grogAPIMessage.Messages = messages
	grogAPIMessage.Model = "Mixtral-8x7b-32768"
	grogAPIMessage.Stream = false
	grogAPIMessage.Temperature = 0
	grogAPIMessage.Reponseformat = map[string]string{"type": "json_object"}

	jsonMessage, err := json.Marshal(grogAPIMessage)
	messageBytes := []byte(jsonMessage)

	client := &http.Client{}
	request, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(messageBytes))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GROQ_API_KEY")))

	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	groqAPIResponse := GroqAPIResponse{}
	err = json.Unmarshal(bodyBytes, &groqAPIResponse)
	if err != nil {
		return "", err
	}

	return groqAPIResponse.Choices[0].Message.Content, nil

}
