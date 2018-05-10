package job

import (
	"github.com/staori/go.uuid"
)

// Bid represents a minimum price or rate a miner is willing to accept to execute a job
type Bid struct {
	ID      uuid.UUID
	JobID   uuid.UUID
	MinerID uuid.UUID
	MinRate float64
}
