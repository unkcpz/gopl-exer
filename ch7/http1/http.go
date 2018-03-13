// http server a database
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"socks": 18, "pants": 50}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}
