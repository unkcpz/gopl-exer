// netcat is a Read-Only tcp client
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8008")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	if _, err := io.Copy(os.Stdout, con); err != nil {
		log.Fatal(err)
	}
}
