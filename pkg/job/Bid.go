package job

// Bid represents a minimum price or rate a miner is willing to accept to execute a job
type Bid struct {
	JobID   string
	MinRate float32
}
