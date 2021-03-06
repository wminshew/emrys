package job

// Specs represents a discrete set of specifications for a job
type Specs struct {
	Rate float64 `json:"rate,omitempty"`
	GPU  string  `json:"gpu,omitempty"`
	RAM  uint64  `json:"ram,omitempty"`
	Disk uint64  `json:"disk,omitempty"`
	Pcie int     `json:"pcie,omitempty"`
}
