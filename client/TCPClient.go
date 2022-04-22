package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

var host2 = flag.String("host", "localhost", "host")
var port2 = flag.String("port", "9091", "port")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host2+":"+*port2)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host2 + ":" + *port2)

	// write
	source := rand.NewSource(time.Now().Unix())
	i := rand.New(source).Intn(100)

	req := "hello " + strconv.Itoa(i)
	_, err = conn.Write([]byte(req))
	if err != nil {
		fmt.Println("error to send message, err:" + err.Error())
		return
	}

	// read
	buf := make([]byte, 512)
	l, err := conn.Read(buf)
	if err != nil {
		fmt.Println("error to receive message, err:" + err.Error())
		return
	}
	fmt.Printf("recieve data %s",buf[:l])
}