package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":8000", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return

	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}

	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	http.ServeFile(w, r, "home.html")

}

func main() {

	flag.Parse()

	hub := newHub()
	go hub.run()

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	fmt.Println(dir)

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome)
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)

	})

	http.Handle("/", r)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)

	}
}
