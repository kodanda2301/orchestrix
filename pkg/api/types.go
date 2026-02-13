package api

type LLMRequest struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
	MaxTokens int `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

type LLMResponse struct {
	Text      string `json:"text"`
	TokenUsage int   `json:"token_usage"`
	Duration  string `json:"duration"`
}

type HealthStatus struct {
	Status    string  `json:"status"`  
	CPUUsage  float64 `json:"cpu_usage"`  
	RAMFree   uint64  `json:"ram_free"` 
	ActiveJob bool    `json:"active_job"`
}