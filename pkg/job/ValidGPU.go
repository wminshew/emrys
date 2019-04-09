package job

import (
	"fmt"
	"strings"
)

const (
	teslaV100     = "tesla_v100"
	teslaP100     = "tesla_p100"
	teslaK80      = "tesla_k80"
	nvidiaTitanV  = "nvidia_titan_v"
	nvidiaTitanXp = "nvidia_titan_xp"
	nvidiaTitanX  = "nvidia_titan_x"
	rtx2080ti     = "geforce_rtx_2080_ti"
	rtx2080       = "geforce_rtx_2080"
	rtx2070       = "geforce_rtx_2070"
	gtx1080ti     = "geforce_gtx_1080_ti"
	gtx1080       = "geforce_gtx_1080"
	gtx1070ti     = "geforce_gtx_1070_ti"
	gtx1070       = "geforce_gtx_1070"
)

// names sourced from https://developer.nvidia.com/cuda-gpus
var validGPU = map[string]string{
	teslaV100:              teslaV100,
	"teslav100":            teslaV100,
	"v100":                 teslaV100,
	teslaP100:              teslaP100,
	"teslap100":            teslaP100,
	"p100":                 teslaP100,
	teslaK80:               teslaK80,
	"teslak80":             teslaK80,
	"k80":                  teslaK80,
	nvidiaTitanV:           nvidiaTitanV,
	"nvidiatitanv":         nvidiaTitanV,
	rtx2080ti:              rtx2080ti,
	"geforcertx2080ti":     rtx2080ti,
	"rtx2080ti":            rtx2080ti,
	"2080ti":               rtx2080ti,
	nvidiaTitanXp:          nvidiaTitanXp,
	"nvidiatitanxp":        nvidiaTitanXp,
	nvidiaTitanX:           nvidiaTitanX,
	"nvidiatitanx":         nvidiaTitanX,
	"nvidiatitanx(pascal)": nvidiaTitanX,
	rtx2080:                rtx2080,
	"geforcertx2080":       rtx2080,
	"rtx2080":              rtx2080,
	"2080":                 rtx2080,
	gtx1080ti:              gtx1080ti,
	"geforcegtx1080ti":     gtx1080ti,
	"gtx1080ti":            gtx1080ti,
	"1080ti":               gtx1080ti,
	rtx2070:                rtx2070,
	"geforcertx2070":       rtx2070,
	"rtx2070":              rtx2070,
	"2070":                 rtx2070,
	gtx1080:                gtx1080,
	"geforcegtx1080":       gtx1080,
	"gtx1080":              gtx1080,
	"1080":                 gtx1080,
	gtx1070ti:              gtx1070ti,
	"geforcegtx1070ti":     gtx1070ti,
	"gtx1070ti":            gtx1070ti,
	"1070ti":               gtx1070ti,
	gtx1070:                gtx1070,
	"geforcegtx1070":       gtx1070,
	"gtx1070":              gtx1070,
	"1070":                 gtx1070,
}

// ValidateGPU returns a well-defined server gpu input, given client input
func ValidateGPU(input string) (string, bool) {
	s := strings.ToLower(strings.Replace(input, " ", "", -1))
	validString, ok := validGPU[s]
	return validString, ok
}

// CompareGPU returns true if gpuA is better than or equal to gpuB
func CompareGPU(gpuA, gpuB string) (bool, error) {
	validA, ok := ValidateGPU(gpuA)
	if !ok {
		return false, fmt.Errorf("invalid gpu: %s", gpuA)
	}
	validB, ok := ValidateGPU(gpuB)
	if !ok {
		return false, fmt.Errorf("invalid gpu: %s", gpuB)
	}
	return compareGPU[validA] >= compareGPU[validB], nil
}

// gpu rankings, loosely based on
// https://browser.geekbench.com/cuda-benchmarks
var compareGPU = map[string]int{
	teslaV100:     100,
	nvidiaTitanV:  90,
	teslaP100:     80,
	rtx2080ti:     75,
	nvidiaTitanXp: 60,
	nvidiaTitanX:  50,
	rtx2080:       45,
	gtx1080ti:     40,
	rtx2070:       35,
	gtx1080:       30,
	gtx1070ti:     20,
	gtx1070:       10,
	teslaK80:      0,
}
