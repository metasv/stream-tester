package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// set go routine(thread) count here
	for i := 0; i < 10; i++ {
		go requestAddress()
		time.Sleep(time.Millisecond * 20)
	}

	select {}

}

// create a go routine to keep connection to the address stream
// if any of the connection is lost, the program panics
func requestAddress() {
	request, _ := http.NewRequest("GET", "https://api.metasv.com/v1/stream/address/1GjJXSYp9DhkvNV8NXfBAecpa5PhtEvpw8", nil)
	// todo replace the api-key
	request.Header.Add("x-api-key", "ask-me-for-the-key")
	request.Header.Set("Content-Type", "application/json")
	clt := http.Client{}
	resp, _ := clt.Do(request)

	reader := bufio.NewReader(resp.Body)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			break
		}
		fmt.Println(string(line))
		if err != nil {
			break
		}
	}
	panic("connection finished")
}
