package main

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const MaxOutstanding = 100
const MaxHashCount = 10000
const (
	InProgress int = 0
	Done           = 1
	NotSeen        = 2
)

type HashRequest struct {
	id     int
	status int
	pwd    string
	value  string
}

var hashes [MaxHashCount]*HashRequest

var id int
var idGenMutex = &sync.Mutex{}

var queue chan *HashRequest
var quit chan bool
var srv *http.Server

func main() {
	http.DefaultServeMux.HandleFunc("/hash/", http.HandlerFunc(hashGetHandler))
	http.DefaultServeMux.HandleFunc("/hash", http.HandlerFunc(hashCreateHandler))
	http.DefaultServeMux.HandleFunc("/shutdown", http.HandlerFunc(shutdownHandler))
	http.DefaultServeMux.HandleFunc("/stats", http.HandlerFunc(statsHandler))

	queue = make(chan *HashRequest, MaxOutstanding)
	quit = make(chan bool)
	go serveHashCreate(queue, quit)

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

	idGenMutex.Lock()
	defer idGenMutex.Unlock()
	id += 1
	hashId := id
	hashes[hashId] = &HashRequest{
		id:     hashId,
		status: InProgress,
		pwd:    inputPwd,
		value:  ""}
	queue <- hashes[hashId]

	w.Write([]byte(fmt.Sprintf("%d", hashId)))
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

	hashResult := hashes[hashId%MaxHashCount] // support < 10k hashes.
	if hashResult.status == Done {
		w.Write([]byte(hashResult.value))
	} else {
		w.WriteHeader(http.StatusNotFound)
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

func serveHashCreate(hashRequests chan *HashRequest, quit chan bool) {
	for i := 0; i < MaxOutstanding; i++ {
		go handleHashCreate(hashRequests)
	}
	<-quit
	close(queue)
}

func handleHashCreate(queue chan *HashRequest) {
	for r := range queue {
		process(r)
	}
}

func process(req *HashRequest) {
	startTime := time.Now()
	time.Sleep(5 * time.Second)
	sha := sha512.New()
	sha.Write([]byte(req.pwd))
	req.value = base64.URLEncoding.EncodeToString(sha.Sum(nil))
	req.status = Done
	hashes[req.id] = req
	go addOP(time.Since(startTime).Milliseconds())
}
