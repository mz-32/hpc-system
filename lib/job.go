package lib

// import (
// 	"encoding/json"
// )

type Job struct {
  Pid string `json:pid`
  Username string `json:username`
  Node string `json:nodeserver`
}
