package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"
)

func main() {
	os.Setenv("URI", "http://localhost:8080/request")
	uri := os.Getenv("URI")
	t := time.Now()
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 100 * time.Second,
	}
	for i := 0; i < 100000; i++ {
		performRequest(uri, client)
	}
	log.Println("Time Taken : ", time.Since(t))
}

func performRequest(uri string, c *http.Client) {

	//	c := httpClient()
	// client trace to log whether the request's underlying tcp connection was re-used
	// c = httpClient()
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { log.Printf("conn was reused: %t", info.Reused) },
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	req, err := http.NewRequestWithContext(traceCtx, "GET", uri, nil)
	if err != nil {
		log.Println("error creating the request", req)
		return
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Println("error performing http request")
		return
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	log.Println("Response : ", string(b))

}

func httpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}

	return client
}
