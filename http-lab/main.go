package main

import (
	"log"
	"net/http"
	"time"
)

var m = map[string]chan bool{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("App start")

	// colors := map[string]string{}

	log.Println("m: ", m)

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/startJob", startJob)
	http.HandleFunc("/stopJob", stopJob)
	http.HandleFunc("/listJobMap", listJobMap)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testJob(name string, stopSig chan bool) {
	for {
		select {
		case <-stopSig:
			log.Printf("Stop.%s", name)
			return
		default:
			log.Printf("Job here %s", name)
			time.Sleep(2 * time.Second)
		}
	}
}
