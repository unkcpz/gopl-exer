// Clock1 is a TCP server write periodic time
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8008")
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleCon(con)
	}
}

func handleCon(con net.Conn) {
	defer con.Close()
	for {
		_, err := io.WriteString(con, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
