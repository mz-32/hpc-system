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
    "math/rand"
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
  go  comandListener()
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
    // get API で各計算ノードからステータス取得
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
func getReqNode(username string, rcpus int, rmem int )int{
  var cpulim = 10.0
  if rcpus <=0{
    rcpus = 2
  }
  if rmem <=0 {
    rmem = 100
  }

  // オンライン判定, 使用中
  var onlineNodes []int
  var limitTime = time.Now().Add(-1* time.Minute)
  for _, node := range config.Server.Nodes {
    _, ok := nodeStat[node.Id]
    if ok  {

      var logind = false
      for _, u:= range nodeStat[node.Id][len(nodeStat[node.Id])-1].ActiveUser{
        if u ==  username{
          logind = true
          break
        }
      }
      t := nodeStat[node.Id][len(nodeStat[node.Id])-1].Time
      statustime, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", t)
      if !statustime.Before(limitTime) || !logind {
        onlineNodes = append(onlineNodes, node.Id)
        }
    }
  }
  // 負荷判定

  var freeNode []int
  for _, nodeId := range onlineNodes {
    var freeCpus = 100
    for _, stat := range nodeStat[nodeId]{
      var c = 0
      for _, cpu := range stat.Cpu {
        cpup := (cpu.User+cpu.System)/(cpu.User+cpu.System+cpu.Idle)
        if cpup < cpulim {
          c++
        }
      }
      if freeCpus > c{
        freeCpus = c
      }
    }
    if freeCpus > rcpus + 2{
      freeNode = append(freeNode,nodeId)
    }
  }

  //　当てはまる中からランダムに計算ノードを選択する
  if len(freeNode) > 0{
    rand.Seed(time.Now().UnixNano())
    i := rand.Intn(len(freeNode))
    return freeNode[i]
  }

  return 0
}

func comandListener(){
  service := ":"+config.Server.SocketPort
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
          fmt.Println(username)
          rcpus, _ :=strconv.Atoi(stra[2])
          rmem, _ :=strconv.Atoi(stra[3])
          id := getReqNode(username, rcpus, rmem)
          if id != 0{
            var ip string
            for _, node := range config.Server.Nodes{
              if node.Id == id{
                ip = node.Ip
                break
              }
            }
            _, err = conn.Write([]byte(ip))
          }else{
            _, err = conn.Write([]byte("err"))
          }


          // freeNode := getFreeNode(cpus,mem)
          // if freeNode.Id==0{
          //   _, err = conn.Write([]byte("err"))
          // }else{
          //   _, err = conn.Write([]byte(freeNode.Ip))
          // }

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
