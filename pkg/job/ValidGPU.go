package job

import (
	"strings"
)

const (
	teslaV100     = "tesla_v100"
	teslaP100     = "tesla_p100"
	teslaK80      = "tesla_k80"
	nvidiaTitanV  = "nvidia_titan_v"
	nvidiaTitanXp = "nvidia_titan_xp"
	nvidiaTitanX  = "nvidia_titan_x"
	gtx1080ti     = "geforce_gtx_1080_ti"
	gtx1080       = "geforce_gtx_1080"
	gtx1070ti     = "geforce_gtx_1070_ti"
	gtx1070       = "geforce_gtx_1070"
)

var validGPU = map[string]string{
	"teslav100":        teslaV100,
	"v100":             teslaV100,
	"teslap100":        teslaP100,
	"p100":             teslaP100,
	"teslak80":         teslaK80,
	"k80":              teslaK80,
	"nvidiatitanv":     nvidiaTitanV,
	"nvidiatitanxp":    nvidiaTitanXp,
	"nvidiatitanx":     nvidiaTitanX,
	"geforcegtx1080ti": gtx1080ti,
	"gtx1080ti":        gtx1080ti,
	"1080ti":           gtx1080ti,
	"geforcegtx1080":   gtx1080,
	"gtx1080":          gtx1080,
	"1080":             gtx1080,
	"geforcegtx1070ti": gtx1070ti,
	"gtx1070ti":        gtx1070ti,
	"1070ti":           gtx1070ti,
	"geforcegtx1070":   gtx1070,
	"gtx1070":          gtx1070,
	"1070":             gtx1070,
}

// ValidateGPU returns a well-defined server gpu input, given client input
func ValidateGPU(input string) (string, bool) {
	s := strings.ToLower(strings.Replace(input, " ", "", -1))
	validString, ok := validGPU[s]
	return validString, ok
}
