package main

import (
	"fmt"
	"net"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// listen
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		// asynchronously respond to request
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	fmt.Println("Talking on a connection")
	for i := 0; i < 20; i++ {
		conn.Write([]byte(fmt.Sprintf("%d\n", i)))
		time.Sleep(200 * time.Millisecond)
	}
	conn.Close()
}
