package job

import (
	"github.com/satori/go.uuid"
)

// BidResp holds a valid JWT after winning bid
type BidResp struct {
	Token string `json:"token"`
}
