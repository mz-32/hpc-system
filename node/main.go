package main
import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
  "time"
  "encoding/json"
  //"github.com/shirou/gopsutil/mem"
  "github.com/shirou/gopsutil/cpu"
)
type StatusJson struct {
    Type     string `json:"type"`
    Data string `json:"data"`
}
func getSutatus()string{
  cpuData, _ := cpu.Percent(time.Duration(5000)*time.Millisecond,false)
  //vmem, _ := mem.VirtualMemory()
  cpuPar := StatusJson{Type:"cpuPar",Data:fmt.Sprintf("%f",cpuData[0])}
  nowTime := StatusJson{Type:"time",Data:time.Now().Format("2006-01-02 15:04:05")}
  data := []StatusJson{nowTime,cpuPar}
  jsonBytes, err := json.Marshal(data)
    if err != nil {
        fmt.Println("JSON Marshal error:", err)
        return ""
    }
    fmt.Println(string(jsonBytes))
    return string(jsonBytes)
}
func newClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)
	client := &Client{
		conn:     connection,
		incoming: make(chan string),
		outgoing: make(chan string),
		reader:   reader,
		writer:   writer,
	}
	go client.read()
	go client.write()
	return client
}
type Client struct {
	conn     net.Conn
	incoming chan string
	outgoing chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
}
func (client *Client) read() {
	for {
		line, err := client.reader.ReadString('\n')
		if err == io.EOF {
			client.conn.Close()
			break
		}
		if err != nil {
			checkError(err, "ReadString Error")
		}
		client.incoming <- line
		fmt.Printf("[%s]Read:%s", client.conn.RemoteAddr(), line)
	}
}
func (client *Client) write() {
	for data := range client.outgoing {
		client.writer.WriteString(getSutatus())
		client.writer.Flush()
		fmt.Printf("[%s]Write:%s\n", client.conn.RemoteAddr(), data)
	}
}
func newListener() *net.TCPListener {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, "Resolve Error")
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "Listen Error")
	fmt.Printf("Server Run Port: %s\n", service)
	return listener
}
func newTCPServer() *Server {
	listener := newListener()
	server := &Server{
		listener: listener,
		clients:  make([]*Client, 0),
		conn:     make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}
	return server
}
type Server struct {
	listener *net.TCPListener
	clients  []*Client
	conn     chan net.Conn
	incoming chan string
	outgoing chan string
}
func (server *Server) acceptLoop() {
	defer server.listener.Close()
	fmt.Println("Ready For Accept")
	for {
		conn, err := server.listener.Accept()
		checkError(err, "Accept Error")
		server.conn <- conn
	}
}
func (server *Server) listen() {
	fmt.Println("Ready For Listen")
	go func() {
		for {
			select {
			case conn := <-server.conn:
				server.addClient(conn)
			case data := <-server.incoming:
				server.response(data)
			}
		}
	}()
}
func (server *Server) addClient(conn net.Conn) {
	fmt.Printf("[%s]Accept\n", conn.RemoteAddr())
	client := newClient(conn)
	server.clients = append(server.clients, client)
	go func() {
		for {
			server.incoming <- <-client.incoming
			client.outgoing <- <-server.outgoing
		}
	}()
}
func (server *Server) response(data string) {
	server.outgoing <- data
}
func checkError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", msg, err.Error())
		os.Exit(1)
	}
}
func main() {
  getSutatus()
	server := newTCPServer()
	server.listen()
	server.acceptLoop()
}
