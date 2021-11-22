package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var srv *http.Server

func main() {
	http.DefaultServeMux.HandleFunc("/hash/", http.HandlerFunc(hashGetHandler))
	http.DefaultServeMux.HandleFunc("/hash", http.HandlerFunc(hashCreateHandler))
	http.DefaultServeMux.HandleFunc("/shutdown", http.HandlerFunc(shutdownHandler))
	http.DefaultServeMux.HandleFunc("/stats", http.HandlerFunc(statsHandler))

	srv = &http.Server{
		Addr: "localhost:8080",
	}
	log.Fatal(srv.ListenAndServe())
}

func hashCreateHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.ParseForm()
	inputPwd := req.Form.Get("password")
	if len(inputPwd) == 0 || len(inputPwd) > 1024 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashId := createHash(inputPwd) // hashId will be return immediately, while the hash process is happening async.
	w.Write([]byte(fmt.Sprintf("%d", hashId)))
}

func hashGetHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	segs := strings.Split(req.URL.Path[1:], "/")
	if len(segs) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var hashId int
	var err error
	if hashId, err = strconv.Atoi(segs[1]); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashResult := getHash(hashId)
	if hashResult.status == Done {
		w.Write([]byte(hashResult.value))
	} else {
		w.WriteHeader(http.StatusNotFound) // return 404 if the hash is still in progress
	}
}

func statsHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	stats := getStats()
	statsJson, _ := json.Marshal(stats)
	w.Write(statsJson)
}

func shutdownHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	quit <- true

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
	close(quit)
}
