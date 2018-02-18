all:
	getlib

getlib:
	go get github.com/shirou/gopsutil
	go get github.com/mattn/go-sqlite3
	go get golang.org/x/sys/unix

run_node:
	go run node/node.go
