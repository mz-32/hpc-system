package lib

import (
	"encoding/json"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"

)

type ServerStat struct {
	// Host
	HostName             string `json:"hostname"`
	HostID               string `json:"hostid"`
	VirtualizationSystem string `json:"virtualizationSystem"`
	// Memory
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	UsedPercent float64 `json:"usedPercent"`
	// DiskIO
	DiskIO []DiskStat `json:"diskIO"`
	// Cpu
	Cpu []cpu.TimesStat `json:"cpu"`
	ActiveUser []string `json;activeUser`
	// Time
	Time string `json:"time"`
	// Error
	ErrorInfo []error `json:"errorInfo"`
}

type DiskStat struct {
	Name       string `json:"name"`
	IoTime     uint64 `json:"ioTime"`
	WeightedIO uint64 `json:"weightedIO"`
}
type Processes struct{
  Pid int32 `json:"pid"`
  Name string `json:"name"`
  Username string `json:"username"`
  IsRunning bool `json:isRunning`
  Nice int32 `json:nice`
  MemoryInfo *process.MemoryInfoStat `json:"memoryInfo"`
}

func (d ServerStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

func (d DiskStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}
func (d Processes) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}
