package main

import (
  "os"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

func generate_mac() {
}

func json_handler(w http.ResponseWriter, r *http.Request) {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	buf[0] |= 2
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"mac\": \"%02x:%02x:%02x:%02x:%02x:%02x\"}\n", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
}

func plain_handler(w http.ResponseWriter, r *http.Request) {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	buf[0] |= 2
	fmt.Fprintf(w, "%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
}

func main() {
	http.HandleFunc("/", plain_handler)
	http.HandleFunc("/plain/", plain_handler)
	http.HandleFunc("/json/", json_handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")),nil))
}
