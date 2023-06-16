package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func ping(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	clientAddr := r.RemoteAddr // <ip>:<port>
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Client IP is %s", strings.Split(clientAddr, ":")[0]) // let's get ip only
	fmt.Fprintf(w, "Hello from %s", hostname)
}

func startJob(w http.ResponseWriter, r *http.Request) {
	sig := make(chan bool)
	id := String(5)
	m[id] = sig
	go testJob(id, sig)
	log.Printf("startJob")
	fmt.Fprintf(w, "startJob")
}

func stopJob(w http.ResponseWriter, r *http.Request) {
	// 	stopSig <- true
	id := r.URL.Query().Get("id")
	log.Println("stopJob:", id)
	m[id] <- true
	fmt.Fprintf(w, "stopJob")
}

func listJobMap(w http.ResponseWriter, r *http.Request) {
	log.Println("listJobMap m: ", m)
	fmt.Fprintf(w, "listJobMap")
}
