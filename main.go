package main

import (
	"net"
	"os"
	"fmt"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:4561")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("MESSAGE\ndestination:fxappender\ncontent-type:text/plain\ncontent-length:2\n\n{}"))
	time.Sleep(3000 * time.Millisecond)
	conn.Close()
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}