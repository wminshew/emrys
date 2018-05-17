package job

// Message represents the structure of messages passed from server to miner
type Message struct {
	Message string `json:"message,omitempty"`
	Job     *Job   `json:"job,omitempty"`
}
