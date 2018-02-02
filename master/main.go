package main

import (
    "fmt"
    "log"
    "net/http"
    "text/template"
    "time"
    "strconv"
    "io/ioutil"
    "encoding/json"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
    //"github.com/shirou/gopsutil/host"
    "./masterdb"

)

type StatusData struct {
  Time time.Time
  Vmem  *mem.VirtualMemoryStat
  CpuPer []float64
}

var a = make(map[int][]StatusData)
var nodes = make(map[int]masterdb.Node)
var onlineNodes = make(map[int]bool)
func main() {
    masterdb.DbInit()
    nodes = masterdb.GetNodeInfo()
    fmt.Println(nodes)

    go handleGetStatus()
    time.Sleep(5 * time.Second)
    http.HandleFunc("/", handleIndex)
    http.HandleFunc("/info", handleInfo)
    http.HandleFunc("/add/", handleAddNode)
    http.HandleFunc("/api/status", handleStatusApi)
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./node_modules/"))))
    // サーバーをポート8080で起動
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
  type IndexData struct {
  Online  bool
  Node masterdb.Node
  }
  var data = make(map[int]IndexData)
  for _, n := range nodes {
    data[n.Id]=IndexData{Online:onlineNodes[n.Id],Node:n}
  }
    t := template.Must(template.ParseFiles("templates/layout.html.tpl"))
    if err := t.ExecuteTemplate(w, "layout.html.tpl",data); err != nil {
        log.Fatal(err)
    }
}
func handleInfo(w http.ResponseWriter, r *http.Request) {
    prams := r.URL.Query()
    _, ok := prams["id"]
    if ok {
      id, _ := strconv.Atoi(prams["id"][0])
      //masterdb.GetNodeInfoFromId(id )
      t := template.Must(template.ParseFiles("templates/info.html.tpl"))

      if err := t.ExecuteTemplate(w, "info.html.tpl",nodes[id]); err != nil {
          log.Fatal(err)
      }
    }else{
      http.Redirect(w, r, "/", http.StatusFound)
    }
}
func handleAddNode(w http.ResponseWriter, r *http.Request) {
    var n masterdb.Node
    prams := r.URL.Query()
    _, ok := prams["id"]
    t := template.Must(template.ParseFiles("templates/addNode.html.tpl"))
    switch r.Method {
    case "POST":
      if ok {
        n.Id, _ = strconv.Atoi(prams["id"][0])
        n.Ip = r.FormValue("IP")
        n.Hostname = r.FormValue("HOSTNAME")
        masterdb.UpdateNode(n)
      }else{
        n.Ip = r.FormValue("IP")
        n.Hostname = r.FormValue("HOSTNAME")
        masterdb.InsertNode(n)
      }
      http.Redirect(w, r, "/", http.StatusFound)
    case "GET":
      if ok {
        pram,_ := strconv.Atoi(prams["id"][0])
        n = masterdb.GetNodeInfoFromId(pram)
      }
      if err := t.ExecuteTemplate(w, "addNode.html.tpl",n ); err != nil {
          log.Fatal(err)
      }
    }
    nodes = masterdb.GetNodeInfo()
}


func handleStatusApi(w http.ResponseWriter, r *http.Request) {
  prams := r.URL.Query()
  _, ok := prams["id"]
  var id int
  if ok {
    id, _ = strconv.Atoi(prams["id"][0])
  } else{
    id =0
  }
  data := a[id][len(a[id])-1]
  outjson, e := json.Marshal(data)
  if e != nil {
      fmt.Println(e)
  }
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprint(w, string(outjson))
}

func handleGetStatus(){
  for ;; {
    go getHostStatus()

    time.Sleep(5000 * time.Millisecond)
  }
}
func getOnlineNodes(){
  t := time.Now().Add(-time.Second*10)
  for _, n := range nodes{
    _, ok := a[n.Id]
    if ok && !t.After(a[n.Id][len(a[n.Id])-1].Time) {
      onlineNodes[n.Id] = true
    }else{
        onlineNodes[n.Id] = false
    }
  }
}
func getNodeStatus(){
  for _, n := range nodes {
    resp, err := http.Get("http://"+n.Ip+":8080/api/status")
  if err != nil {
    fmt.Println(err)
  }else{
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

     if err != nil {
         fmt.Println(err)
     }
     var data StatusData
    if err := json.Unmarshal(body, &data); err != nil {
        log.Fatal(err)
    }
    _, ok := a[n.Id]
    if ok  {
      a[n.Id] = append(a[n.Id], data)
    }else{
      a[n.Id] = []StatusData{data}
    }
    if len(a[n.Id]) > 100 {
      a[n.Id] = a[n.Id][1:]
    }
  }
    }
}
func getHostStatus(){
    fmt.Println("Host status is getting")
    loc, _ := time.LoadLocation("Asia/Tokyo")
    cpuPer, _ := cpu.Percent(time.Duration(500)*time.Millisecond,true)
    nowTime := time.Now().In(loc)
    vmem, _ := mem.VirtualMemory()
    data := StatusData{Time:nowTime,Vmem: vmem, CpuPer: cpuPer}
    _, ok := a[0]
    if ok  {
      a[0] = append(a[0], data)
    }else{
      a[0] = []StatusData{data}
    }
    if len(a) > 100 {
      a[0] = a[0][1:]
    }
    fmt.Println("time : "+nowTime.Format("2006-01-02 15:04:05"))
    getNodeStatus()
    getOnlineNodes()
}
