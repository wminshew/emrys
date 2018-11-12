package job

import (
	"github.com/satori/go.uuid"
)

// Job represents a discrete set of computations submitted by a user to be executed by a miner
type Job struct {
	ID       uuid.UUID `json:"id,omitempty"`
	UserID   uuid.UUID `json:"userId,omitempty"`
	WinBidID uuid.UUID `json:"winBidId,omitempty"`
	Rate     float64   `json:"rate,omitempty"`
}
