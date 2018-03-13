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
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			msg := fmt.Sprintf("no such page: %s\n", r.URL)
			http.Error(w, msg, http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		msg := fmt.Sprintf("no such page: %s\n", r.URL)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
}

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}
