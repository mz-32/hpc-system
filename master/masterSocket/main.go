package masterSocket
import (
    "fmt"
    "net"
    "time"
)

func ask(host string,mes string)(string){
  conn, err := net.Dial("tcp", "127.0.0.1:7777")
  if err != nil {
      fmt.Printf("Dial error: %s\n", err)
      return nil
  }
  defer conn.Close()
  conn.Write([]byte(mes))

  readBuf := make([]byte, 1024)
  conn.SetReadDeadline(time.Now().Add(10 * time.Second))
  readlen, err := conn.Read(readBuf)
  if err != nil {
      fmt.Printf("Dial error: %s\n", err)
      return nil
  }
  fmt.Println("server: " + string(readBuf[:readlen]))

  return string(readBuf[:readlen])
}
