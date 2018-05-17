package job

import (
	"github.com/satori/go.uuid"
)

// Job represents a discrete set of computations submitted by a user to be executed by a miner
type Job struct {
	ID       uuid.UUID `json:"id,omitempty"`
	UserID   uuid.UUID `json:"userId,omitempty"`
	WinBidID uuid.UUID `json:"winBidId,omitempty"`
	PayRate  float64   `json:"payRate,omitempty"`
	Token    string    `json:"token,omitempty"`
}
