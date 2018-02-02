package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "encoding/json"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"

)

type StatusData struct {
  Time time.Time
  Vmem  *mem.VirtualMemoryStat
  CpuPer []float64
}
var a StatusData

func main() {

    go handleGetStatus()
    time.Sleep(5 * time.Second)
    http.HandleFunc("/api/status", handleStatusApi)
    // サーバーをポート8080で起動
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleStatusApi(rw http.ResponseWriter, req *http.Request) {
    outjson, e := json.Marshal(a)
    if e != nil {
        fmt.Println(e)
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
    fmt.Println("Host status is getting")
    loc, _ := time.LoadLocation("Asia/Tokyo")
    cpuPer, _ := cpu.Percent(time.Duration(500)*time.Millisecond,true)
    nowTime := time.Now().In(loc)
    vmem, _ := mem.VirtualMemory()
    a = StatusData{Time:nowTime,Vmem: vmem, CpuPer: cpuPer}
    fmt.Println(a)

}
