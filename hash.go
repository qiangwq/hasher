package main

import (
	"crypto/sha512"
	"encoding/base64"
	"sync"
	"time"
)

const (
	InProgress int = 0 // new HashRequest is InProgress state by default
	Done           = 1 // the hash is ready

	MaxOutstanding = 100   // maxium # of concurrent hash goroutines
	MaxHashCount   = 10000 // support up to MaxHashCount hash in mem
)

var id int                     // current id
var idGenMutex = &sync.Mutex{} // lock for uniquely increasing id

type HashRequest struct {
	id     int    // hash ID
	status int    // InProgress or Done
	pwd    string // input password
	value  string // hash of the password
}

var hashCache [MaxHashCount]*HashRequest // cache of all HashRequests so far

var queue chan *HashRequest // HashRequest queue
var quit chan bool          // quit channel for server to command to close all hash handlers

func init() {
	queue = make(chan *HashRequest, MaxOutstanding)
	quit = make(chan bool)

	go startHashHandlers()
}

func createHash(pwd string) int {
	idGenMutex.Lock()
	defer idGenMutex.Unlock()
	id += 1
	hashCache[id%MaxHashCount] = &HashRequest{ // only cash MaxHashCount amount of hashes
		id:     id,
		status: InProgress,
		pwd:    pwd,
		value:  ""}
	queue <- hashCache[id]
	return id
}

func getHash(id int) *HashRequest {
	return hashCache[id%MaxHashCount] // only cash MaxHashCount amount of hashes
}

func startHashHandlers() {
	for i := 0; i < MaxOutstanding; i++ {
		go handler()
	}
	<-quit
	close(queue)
}

func stopHashHandler() {
	quit <- true
}

func handler() {
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
	hashCache[req.id] = req
	go addOP(time.Since(startTime).Milliseconds())
}
