package main

import (
    "fmt"
    "log"
    "net/http"
    "text/template"
    "time"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
    //"github.com/shirou/gopsutil/host"

)

type StatusData struct {
  Time string
  Vmem  *mem.VirtualMemoryStat
  CpuPer []float64
}
var a []StatusData


func main() {



    go handleGetStatus()
    http.HandleFunc("/", handleIndex)
    //http.HandleFunc("/", handleMem)

    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./node_modules/"))))
    // サーバーをポート8080で起動
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("templates/layout.html.tpl"))

    if err := t.ExecuteTemplate(w, "layout.html.tpl",a[len(a)-1]); err != nil {
        log.Fatal(err)
    }
}
func handleGetStatus(){
  for ;; {
    go getHostStatus()
    time.Sleep(5000 * time.Millisecond)
  }
}
func getHostStatus(){
    fmt.Println("Host status is getting")
    cpuPer, _ := cpu.Percent(time.Duration(5000)*time.Millisecond,false)
    nowTime := time.Now().Format("2006-01-02 15:04:05")
    vmem, _ := mem.VirtualMemory()
    data := StatusData{Time:nowTime,Vmem: vmem, CpuPer: cpuPer}
    a = append(a, data)
    if len(a) > 100 {
      a = a[1:]
    }
    fmt.Println("time : "+nowTime)
    fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", vmem.Total, vmem.Free, vmem.UsedPercent)
    fmt.Println(cpuPer)
    fmt.Println(len(a))
}
