package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "time"
    "encoding/json"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/host"
)

type StatusData struct {
  Time time.Time
  Vmem  *mem.VirtualMemoryStat
  CpuPer []float64
}
var status StatusData
var hostInfo *host.InfoStat
var wlog *log.Logger
var cpuPer []float64
var vmem *mem.VirtualMemoryStat

func init(){
  log.SetPrefix("[Info]")  // 接頭辞の設定
  wlog = log.New(os.Stdout, "[Warning]", log.LstdFlags|log.LUTC)
}
func main() {
    go handleGetStatus()
    hostInfo, _ = host.Info()
    time.Sleep(5 * time.Second)
    log.Print("Start Node")
    http.HandleFunc("/api/status.json", handleStatusApi)
    http.HandleFunc("/api/info.json", handleInfoApi)
    http.HandleFunc("/api/mem/status.json", handleMemStatusApi)
    http.HandleFunc("/api/cpu/per.json", handleCpuPerApi)
    http.HandleFunc("/api/cpu/info.json", handleCpuInfoApi)
    http.HandleFunc("/api/cpu/counts.json", handleCpuCountsApi)

    log.Fatal(http.ListenAndServe(":55010", nil))
}
func handleMemStatusApi(rw http.ResponseWriter, req *http.Request) {
    log.Print("acsessd /api/mem/status")
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, vmem)
}
func handleCpuPerApi(rw http.ResponseWriter, req *http.Request) {
    log.Print("acsessd /api/cpu/per")
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, cpuPer)
}
func handleCpuInfoApi(rw http.ResponseWriter, req *http.Request) {
    log.Print("acsessd /api/cpu/info")
    cpuInfos, _ := cpu.Info()
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, cpuInfos)
}
func handleCpuCountsApi(rw http.ResponseWriter, req *http.Request) {
    log.Print("acsessd /api/cpu/counts")
    cpuCount, _ := cpu.Counts(true)
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, cpuCount)
}
func handleInfoApi(rw http.ResponseWriter, req *http.Request) {
    log.Print("acsessd /api/info")
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, hostInfo)
}
func handleStatusApi(rw http.ResponseWriter, req *http.Request) {
    log.Print("acsessd /api/status")
    outjson, e := json.Marshal(status)
    if e != nil {
      wlog.Print(e)
    }
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, string(outjson))
}

func handleGetStatus(){
  for ;; {
    go getHostStatus()
    time.Sleep(5000 * time.Millisecond)
  }
}
func getHostStatus(){
    loc, _ := time.LoadLocation("Asia/Tokyo")
    cpuPer, _ = cpu.Percent(time.Duration(500)*time.Millisecond,true)
    nowTime := time.Now().In(loc)
    vmem, _ = mem.VirtualMemory()
    status = StatusData{Time:nowTime,Vmem: vmem, CpuPer: cpuPer}
}
