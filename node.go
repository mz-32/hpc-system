package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "./lib"
)

var wlog *log.Logger

func init(){
    log.SetPrefix("[Info]")  // 接頭辞の設定
    wlog = log.New(os.Stdout, "[Warning]", log.LstdFlags|log.LUTC)

    logfile, err := os.OpenFile("./test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
      panic("cannnot open test.log:" + err.Error())
    }
    defer logfile.Close()
    log.SetOutput(logfile)
}
func main() {

    log.Print("Start Node")
    http.HandleFunc("/api/status.json", handleStatusApi)
    log.Fatal(http.ListenAndServe(":55010", nil))
}

func handleStatusApi(rw http.ResponseWriter, req *http.Request) {
    log.Print("acsessd /api/status.json")
    status, e :=lib.GetServerStat()
    if e != nil {
      wlog.Print(e)
    }
    log.Print(status )
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, status)
}
