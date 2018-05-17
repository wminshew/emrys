package job

import (
	"github.com/satori/go.uuid"
)

// Job represents a discrete set of computations submitted by a user to be executed by a miner
type Job struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	WinBidID uuid.UUID
	PayRate  float64
	Token    string
}
