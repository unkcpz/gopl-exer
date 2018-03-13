// http serve database
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"socks": 18, "shoes": 50}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {

}

func (db database) price(w http.ResponseWriter, r *http.Request) {

}

func (db database) update(w http.ResponseWriter, r *http.Request) {

}
