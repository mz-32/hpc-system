all: getlib	node master

getlib:
	go get github.com/shirou/gopsutil
	go get github.com/mattn/go-sqlite3
	go get golang.org/x/sys/unix
	go get github.com/BurntSushi/toml

run_node:
	go run node/node.go

init:
	mkdir ./bin/

node: node.go init
	go build -o ./bin/node node.go

master: master.go
	go build -o ./bin/master master.go
clean:
	rm -r ./bin/
