package main

import (
    "fmt"
    "encoding/json"
    "github.com/shirou/gopsutil/process"
)

type Processes struct{
  Pid int32 `json:"pid"`
  Name string `json:"name"`
  Username string `json:"username"`
  IsRunning bool `json:isRunning`
  Nice int32 `json:nice`
  MemoryInfo *process.MemoryInfoStat `json:"memoryInfo"`
}
func (p Processes) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}


func main(){
  var activeUser []string
  p, e := process.Processes()
  if e != nil {
    fmt.Println(e)
  }
  for _, i := range p {
    u, e := i.Username()
    if e != nil {
      fmt.Println(e)
    }
    var f = true
    for _, au := range activeUser{
      if u == au{
        f = false
        break
      }
    }
    if f  {
      activeUser = append(activeUser,u)
    }
  }
  // var jobs = make(map[string][]Processes)
  // var activeUser []string
  // p, e := process.Processes()
  // if e != nil {
  //   fmt.Println(e)
  // }
  // var pppp Processes
  // for _, i := range p {
  //   u, e := i.Username()
  //   if e != nil {
  //     fmt.Println(e)
  //   }
  // //  if u != "root" {
  //     n, _ := i.Name()
  //     b, _ := i .IsRunning()
  //     nice, _ := i.Nice()
  //     mem, _ := i.MemoryInfo()
  //     pppp =Processes{Pid:i.Pid,Name:n,Username:u,IsRunning:b,Nice:nice,MemoryInfo:mem}
  //     _, ok := jobs[u]
  //     if ok  {
  //       jobs[u] = append(jobs[u], pppp)
  //     }else{
  //       jobs[u] = []Processes{pppp}
  //     }
  //   //  fmt.Println(pppp)
  //   }
  // //}
  // for i, _ := range jobs {
  //   activeUser = append(activeUser,i)
  //   //fmt.Println(i)
  // }
  fmt.Println(activeUser)
}
