package job

import (
	"github.com/satori/go.uuid"
)

// Bid represents a minimum price or rate a miner is willing to accept to execute a job
type Bid struct {
	ID      uuid.UUID `json:"id,omitempty"`
	JobID   uuid.UUID `json:"jobId,omitempty"`
	MinerID uuid.UUID `json:"minerId,omitempty"`
	MinRate float64   `json:"minRate,omitempty"`
}
