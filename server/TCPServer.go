package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var host2 = flag.String("host", "", "host")
var port2 = flag.String("port", "9091", "port")

func main() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host2+":"+*port2)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + *host2 + ":" + *port2)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}

		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	/*
	for {
		io.Copy(conn, conn)
	}
	*/

	// read from conn
	buf := make([]byte, 1024)
	l, err := conn.Read(buf)
	if err != nil {
		fmt.Println("err:" + err.Error())
		return
	}
	fmt.Printf("Received data: %v, len=%d", string(buf[:l]), l)

	// write data to conn
	data := "Sent:" + string(buf[:l])
	fmt.Printf("\nSend data: %v, len=%d", data, len(data))
	_, err = conn.Write([]byte(data))
	if err != nil {
		fmt.Println("err:" + err.Error())
	}
}