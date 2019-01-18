package main

import (
	"flag"
	"fmt"
	"github.com/didil/gcf-go-image-resizer"
	"log"
	"net/http"
)

// server for local testing
func main() {
	port := flag.Int("p",8080,"server port")
	mux := http.NewServeMux()
	mux.HandleFunc("/ResizeImage", gcf_go_image_resizer.ResizeImage)
	fmt.Printf("Starting local server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
