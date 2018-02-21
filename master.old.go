package main

import (
    "os"
    "fmt"
    "log"
    "net"
    "strings"
    "net/http"
    "text/template"
    "time"
    "math/rand"
    "strconv"
    "io/ioutil"
    "encoding/json"
    "github.com/shirou/gopsutil/mem"
    // "github.com/shirou/gopsutil/cpu"
    //"github.com/shirou/gopsutil/host"
    "./masterdb"
    "../lib"

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
    go comandListener()
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
    log.Fatal(http.ListenAndServe(":10080", nil))
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
  type InfoData struct {
  Id  int
  Hostname string
  CpuCount []int
  }

    prams := r.URL.Query()
    _, ok := prams["id"]
    if ok {
      id, _ := strconv.Atoi(prams["id"][0])
      cpus :=getNodeCpuCpunt(id)
      data := InfoData{Id:id,Hostname:nodes[id].Ip,CpuCount:make([]int, cpus)}
      //masterdb.GetNodeInfoFromId(id )
      t := template.Must(template.ParseFiles("templates/info.html.tpl"))

      if err := t.ExecuteTemplate(w, "info.html.tpl",data); err != nil {
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

        masterdb.UpdateNode(n)
      }else{
        n.Ip = r.FormValue("IP")

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
func getNodeCpuCpunt(n int)int{

    resp, err := http.Get("http://"+nodes[n].Ip+":55010/api/cpu/counts.json")
  if err != nil {
    fmt.Println(err)
  }else{
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
     if err != nil {
         fmt.Println(err)
     }
     var data int
    if err := json.Unmarshal(body, &data); err != nil {
        log.Fatal(err)
    }
    return data
  }
  return 0
}
func getNodeStatus(){
  for _, n := range nodes {
    resp, err := http.Get("http://"+n.Ip+":55010/api/status.json")
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
func getFreeNode(rcpu int, rmem int)(masterdb.Node){
  var freeNode masterdb.Node
  var freeNodes []masterdb.Node
  //freecpus := 0

  if rcpu <=0{
    rcpu = 2
  }
  if rmem <=0 {
    rmem = 10
  }
  for i, n:= range onlineNodes{
    if n {
      c := 0
      for _,cpup := range a[i][len(a[i])-1].CpuPer{
            if cpup < 5{
              c++
            }
          }
          if c > rcpu{
            freeNodes = append(freeNodes, nodes[i])
          }

    }
  }
  if len(freeNodes) >0{
    freeNode = freeNodes[rand.Intn(len(freeNodes))]
  }
  return freeNode

}
func getHostStatus(){
    getNodeStatus()
    getOnlineNodes()
    //fmt.Println(getFreeNode())
}
func comandListener(){
  service := ":10081"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    chkErr(err, "ResolveTCPAddr")
    listener, err := net.ListenTCP("tcp", tcpAddr)
    chkErr(err, "ListenTCP")
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        buf := make([]byte, 1024)
        n, _ := conn.Read(buf)
        //chkErr(err, "Read")
        str :=string(buf[:n])
        stra := strings.Split(str, ",")
        input, _ :=strconv.Atoi(stra[0])
        fmt.Println(stra)
        if input == 0 {
          cpus, _ :=strconv.Atoi(stra[1])
          mem, _ :=strconv.Atoi(stra[2])
          freeNode := getFreeNode(cpus,mem)
          if freeNode.Id==0{
            _, err = conn.Write([]byte("err"))
          }else{
            _, err = conn.Write([]byte(freeNode.Ip))
          }
          _ = conn.Close()
        }else{
          _, err = conn.Write([]byte(""))
          chkErr(err, "Write")
          _ = conn.Close()
        }

        //node := getFreeNode()
        //daytime := time.Now().String()


    }
}

func chkErr(err error, place string) {
    if err != nil {
        fmt.Printf("(%s)", place)
        fmt.Fprintf(os.Stderr, "%s", err.Error())
        os.Exit(0)
    }
}
