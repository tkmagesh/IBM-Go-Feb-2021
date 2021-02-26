package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	responseCh := make(chan *http.Response)
	go getResponse(responseCh)
	for {
		select {
		case res := <-responseCh:
			io.Copy(os.Stdout, res.Body)
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Println("timeout!!")
			return
		}
	}
}

func getResponse(ch chan *http.Response) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- resp
}
