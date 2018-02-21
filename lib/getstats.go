package lib

import (
	"fmt"
	// "os/exec"
	// "strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

func GetServerStat() (ServerStat, []error) {
	var d ServerStat
	var err []error

	errHost := d.GetHostStat()
	if errHost != nil {
		err = append(err, errHost)
	}
	errMemory := d.GetMemoryStat()
	if errMemory != nil {
		err = append(err, errMemory)
	}
	errDisk := d.GetDiskIOStat()
	if errDisk != nil {
		err = append(err, errDisk)
	}
	errCpu := d.GetCpuStat()
	if errCpu != nil {
		err = append(err, errCpu)
	}
	d.GetTime()
	errAU := d.GetActiveUsers()
	if errAU != nil {
		err = append(err, errAU)
	}
	if err != nil {
		return d, err
	}

	return d, nil
}

func (s *ServerStat) GetHostStat() error {
	h, err := host.Info()
	if err != nil {
		return err
	}
	s.HostName = h.Hostname
	s.HostID = h.HostID
	s.VirtualizationSystem = h.VirtualizationSystem
	return nil
}

func (s *ServerStat) GetMemoryStat() error {
	m, err := mem.VirtualMemory()
	if err != nil {
		return err
	}
	s.Total = m.Total
	s.Available = m.Available
	s.UsedPercent = m.UsedPercent
	return nil
}

func (s *ServerStat) GetDiskIOStat() error {
	var ds []DiskStat
	i, err := disk.IOCounters()
	if err != nil {
		return err
	}
	for k, v := range i {
		var d DiskStat
		d.Name = k
		d.IoTime = v.IoTime
		d.WeightedIO = v.WeightedIO
		ds = append(ds, d)
	}
	s.DiskIO = ds
	return nil
}

func (s *ServerStat) GetTime() {
	now := time.Now()
	s.Time = fmt.Sprint(now)
}

func (s *ServerStat) GetCpuStat() error {
	c, err := cpu.Times(true)
	if err != nil {
		return err
	}
	s.Cpu = c
	return nil
}

func (s *ServerStat)GetActiveUsers() error{
  var users []string
  p, e := process.Processes()
  if e != nil {
    return e
  }
  for _, i := range p {
    u, e := i.Username()
    if e != nil {
      return e
    }
		var f = true
		for _, au := range users{
			if u == au{
				f = false
				break
			}
		}
		if f {
			users = append(users,u)
		}
  }
	s.ActiveUser = users
	return nil
}
