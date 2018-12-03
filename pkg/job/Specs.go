package job

// Specs represents a discrete set of specifications for a job
type Specs struct {
	Rate    float64 `json:"rate,omitempty"`
	GPU     string  `json:"gpu,omitempty"`
	RAM     uint64  `json:"ram,omitempty"`
	ramStr  string
	Disk    uint64 `json:"disk,omitempty"`
	diskStr string
	Pcie    int `json:"pcie,omitempty"`
	pcieStr string
}
