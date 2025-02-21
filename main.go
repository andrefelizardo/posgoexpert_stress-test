package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
    url := flag.String("url", "", "URL do serviço a ser testado")
    totalRequests := flag.Int("requests", 1, "Número total de requests")
    concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
    flag.Parse()

    start := time.Now()
    var wg sync.WaitGroup
    requestsCh := make(chan int, *totalRequests)
    statusCount := make(map[int]int)
    var mu sync.Mutex

    for i := 0; i < *totalRequests; i++ {
        requestsCh <- i
    }
    close(requestsCh)

    for i := 0; i < *concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for range requestsCh {
                resp, err := http.Get(*url)
                if err == nil {
                    mu.Lock()
                    statusCount[resp.StatusCode]++
                    mu.Unlock()
                    resp.Body.Close()
                }
            }
        }()
    }

    wg.Wait()
    elapsed := time.Since(start)

    totalOK := statusCount[200]
    fmt.Printf("Tempo total (s): %.2f\n", elapsed.Seconds())
    fmt.Printf("Total requests: %d\n", *totalRequests)
    fmt.Printf("Requests com status 200: %d\n", totalOK)
    fmt.Println("Distribuição de status:")
    for code, count := range statusCount {
        fmt.Printf("  %d: %d\n", code, count)
    }
}