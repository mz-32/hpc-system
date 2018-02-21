package main

import (
    "os"
    "fmt"
    "time"
    "log"
    "strconv"
    "strings"
    "net"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "./lib"
)

var wlog *log.Logger
var config lib.Config
var nodeStat = make(map[int][]lib.ServerStat)
var jobs = make(map[int][]lib.Job)

func init(){
    log.SetPrefix("[Info]")  // 接頭辞の設定
    wlog = log.New(os.Stdout, "[Warning]", log.LstdFlags|log.LUTC)

    var errors []error
    config, errors = lib.GetServerConfig()
    if errors != nil {
      fmt.Println(errors)
    }

}
func main (){

  //go updateData()
  fmt.Println(len(config.Server.Nodes))
  go updateData()
  http.HandleFunc("/", handleIndex)
  log.Fatal(http.ListenAndServe(":"+config.Server.WebPort, nil))

}


func handleIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}

func updateData(){
  i := time.Duration(config.Server.UpdateInterval)
  for ;; {
    go getNodeStats()
    time.Sleep(i * time.Millisecond)
  }
}

func getNodeStats(){
  for _, node := range config.Server.Nodes {

    resp, err := http.Get("http://"+node.Ip+":"+node.Port+"/api/status.json")
    if err != nil {
      fmt.Println(err)
    }else{
      defer resp.Body.Close()
      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
          fmt.Println(err)
      }
      var data lib.ServerStat
      if err := json.Unmarshal(body, &data); err != nil {
          log.Fatal(err)
      }
      _, ok := nodeStat[node.Id]
      if ok  {
        nodeStat[node.Id] = append(nodeStat[node.Id], data)
      }else{
        nodeStat[node.Id] = []lib.ServerStat{data}
      }
      if len(nodeStat[node.Id]) > 100 {
        nodeStat[node.Id] = nodeStat[node.Id][1:]
      }
    }
  }
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
        mode, _ :=strconv.Atoi(stra[0])
        username := stra[1]
        fmt.Println(stra)
        if mode == 0 {
        
          // freeNode := getFreeNode(cpus,mem)
          // if freeNode.Id==0{
          //   _, err = conn.Write([]byte("err"))
          // }else{
          //   _, err = conn.Write([]byte(freeNode.Ip))
          // }
          _, err = conn.Write([]byte(username))
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
