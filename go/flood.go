package main

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	targetURL = "http://localhost:8080"
	requests  = 1000
	workers   = 10
)

func flood(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < requests/workers; i++ {
		resp, err := http.Get(targetURL)
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}
		resp.Body.Close()
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go flood(&wg)
	}
	wg.Wait()
}
