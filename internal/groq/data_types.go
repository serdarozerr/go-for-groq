package groq

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqAPIQuery struct {
	Model         string            `json:"model"`
	Reponseformat map[string]string `json:"response_format"`
	Stream        bool              `json:"stream"`
	Temperature   int               `json:"temperature"`
	Messages      []Message         `json:"messages"`
}

type GroqAPIResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		QueueTime        float64 `json:"queue_time"`
		PromptTokens     int     `json:"prompt_tokens"`
		PromptTime       float64 `json:"prompt_time"`
		CompletionTokens int     `json:"completion_tokens"`
		CompletionTime   float64 `json:"completion_time"`
		TotalTokens      int     `json:"total_tokens"`
		TotalTime        float64 `json:"total_time"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
	XGroq             struct {
		Id string `json:"id"`
	} `json:"x_groq"`
}

type GroqUserQuery struct {
	Query  string `json:"query"`
	UserId string `json:"user_id"`
}
type GroqUserResponse struct {
	Answer string `json:"answer"`
}
