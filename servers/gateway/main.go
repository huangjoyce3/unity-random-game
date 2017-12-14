package main

import (
	"fmt"
	"github.com/huangjoyce3/unity/servers/gateway/handlers"
	"log"
	"net/http"
	"os"
)

const (
	summaryPath = "/summary"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlskey := os.Getenv("TLSKEY")
	tlscert := os.Getenv("TLSCERT")
	if len(tlskey) == 0 || len(tlscert) == 0 {
		log.Fatal("please set TLSKEY and TLSCERT")
	}

	mux := http.NewServeMux()
	mux.HandleFunc(summaryPath, handlers.GameSummaryHandler)
	corsHandler := handlers.NewCORSHandler(mux)

	fmt.Printf("server is listening at http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, corsHandler))
}
