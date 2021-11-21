package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func logErr(s string, err error) {
	fmt.Printf("%s, error %s\n", s, err)
}

func logInfo(s string) {
	fmt.Printf("Info: %s\n", s)
}

func main() {
	quit := make(chan bool)
	sem := make(chan int, 2)
	for i := 0; i < 100; i++ {
		sem <- 1
		go func(i int) {
			if err := test(i); err != nil {
				logErr("test failed with", err)
			}
			time.Sleep(2 * time.Second)
			<-sem
		}(i)
	}
	<-quit
}

func test(tester int) (err error) {
	baseUrl := "http://localhost:8080"
	hashUrl := fmt.Sprintf("%s/%s", baseUrl, "hash")
	statsUrl := fmt.Sprintf("%s/%s", baseUrl, "stats")

	var resp1 *http.Response
	if resp1, err = http.PostForm(hashUrl, url.Values{"password": {"abc"}}); err != nil {
		return
	}

	defer resp1.Body.Close()
	id, err := io.ReadAll(resp1.Body)
	if err != nil {
		return
	}
	logInfo(fmt.Sprintf("tester: %d, %s returns hashID %s", tester, hashUrl, id))

	for {
		urlStr := fmt.Sprintf("%s/%s", hashUrl, id)
		var resp2 *http.Response
		if resp2, err = http.Get(urlStr); err != nil {
			return
		}

		if resp2.StatusCode == 200 {
			defer resp2.Body.Close()
			var hashValue []byte
			if hashValue, err = io.ReadAll(resp2.Body); err != nil {
				return
			}
			logInfo(fmt.Sprintf("tester: %d, hashID: %s, hashValue: %s", tester, id, hashValue))

			if resp2, err = http.Get(statsUrl); err == nil {
				var stats []byte
				if stats, err = io.ReadAll(resp2.Body); err == nil {
					logInfo(fmt.Sprintf("stats: %s", stats))
				}
			}
			return
		} else if resp2.StatusCode == 404 {
			logInfo(fmt.Sprintf("tester: %d, hashID %s not ready yet", tester, id))
			time.Sleep(2 * time.Second)
		}
	}
}
