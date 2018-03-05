// TCP server echo time.Now() periodiclly
// port needed
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8008, "clock lisen port")

func main() {
	flag.Parse()
	url := fmt.Sprintf("localhost:%d", *port)
	l, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleCon(conn)
	}
}

func handleCon(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
