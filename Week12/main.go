package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	dir := flag.String("dir", ".", "The directory of static files to host (default: current directory)")
	port := flag.Int("port", 8080, "The port to listen on (default: 8080)")
	flag.Parse()

	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", *dir)
	}

	fs := http.FileServer(http.Dir(*dir))

	http.Handle("/", fs)

	address := fmt.Sprintf(":%d", *port)
	log.Printf("Serving %s on HTTP port %d\n", *dir, *port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
