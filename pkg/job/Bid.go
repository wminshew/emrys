package job

import (
	"github.com/satori/go.uuid"
)

// Bid represents a minimum price or rate a miner is willing to accept to execute a job on a device
type Bid struct {
	ID       uuid.UUID `json:"id,omitempty"`
	JobID    uuid.UUID `json:"jobId,omitempty"`
	MinerID  uuid.UUID `json:"minerId,omitempty"`
	DeviceID uuid.UUID `json:"deviceId,omitempty"`
	BidRate  float64   `json:"BidRate,omitempty"`
	Late     bool      `json:"late,omitempty"`
}
