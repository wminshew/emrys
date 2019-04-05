package job

import (
	"docker.io/go-docker/api/types"
	"github.com/satori/go.uuid"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// MinerStats represent the state of a miner to be monitored by the server
type MinerStats struct {
	CPUInfo     []cpu.InfoStat         `json:"cpu_info,omitempty"`
	CPUTimes    []cpu.TimesStat        `json:"cpu_times,omitempty"`
	Mem         *mem.VirtualMemoryStat `json:"mem,omitempty"`
	Disk        *disk.UsageStat        `json:"disk,omitempty"`
	WorkerStats []*WorkerStats         `json:"worker_stats,omitempty"`
}

// WorkerStats represent the state of a worker to be monitored by the server
type WorkerStats struct {
	JobID       uuid.UUID       `json:"job_id,omitempty"`
	GPUStats    *DeviceSnapshot `json:"gpu_stats,omitempty"`
	DockerStats *types.Stats    `json:"docker_stats,omitempty"`
	DockerDisk  *DockerDisk     `json:"docker_disk,omitempty"`
}

// DockerDisk represents the disk usage of a worker to be monitored by the server
type DockerDisk struct {
	SizeRw        int64  `json:"size_rw,omitempty"`
	SizeRootFs    int64  `json:"size_root_fs,omitempty"`
	SizeDataDir   uint64 `json:"data_folder,omitempty"`
	SizeOutputDir uint64 `json:"output_folder,omitempty"`
}

// DeviceSnapshot holds data collected about the mining GPU
type DeviceSnapshot struct {
	ID                uuid.UUID `json:"id,omitempty"`
	TimeStamp         int64     `json:"timestamp,omitempty"`
	Active            bool      `json:"busy,omitempty"`
	MinorNumber       uint      `json:"minornumber,omitempty"`
	Name              string    `json:"name,omitempty"`
	Brand             uint      `json:"brand,omitempty"`
	PersistenceMode   uint      `json:"persistence_mode,omitempty"`
	ComputeMode       uint      `json:"compute_mode,omitempty"`
	PerformanceState  uint      `json:"pstate,omitempty"`
	AvgGPUUtilization uint      `json:"avg_gpu_util,omitempty"`
	AvgPowerUsage     uint      `json:"avg_power,omitempty"`
	PowerLimit        uint      `json:"power_limit,omitempty"`
	DefaultPowerLimit uint      `json:"default_power_limit,omitempty"`
	TotalMemory       uint64    `json:"total_mem,omitempty"`
	UsedMemory        uint64    `json:"used_mem,omitempty"`
	GrClock           uint      `json:"gr_clock,omitempty"`
	SMClock           uint      `json:"sm_clock,omitempty"`
	MemClock          uint      `json:"mem_clock,omitempty"`
	GrMaxClock        uint      `json:"gr_max_clock,omitempty"`
	SMMaxClock        uint      `json:"sm_max_clock,omitempty"`
	MemMaxClock       uint      `json:"mem_max_clock,omitempty"`
	PcieTxThroughput  uint      `json:"pcie_tx_thru,omitempty"`
	PcieRxThroughput  uint      `json:"pcie_rx_thru,omitempty"`
	PcieGeneration    uint      `json:"pcie_gen,omitempty"`
	PcieWidth         uint      `json:"pcie_width,omitempty"`
	PcieMaxGeneration uint      `json:"pcie_max_gen,omitempty"`
	PcieMaxWidth      uint      `json:"pcie_max_width,omitempty"`
	Temperature       uint      `json:"temp,omitempty"`
	FanSpeed          uint      `json:"fan_speed,omitempty"`
}
