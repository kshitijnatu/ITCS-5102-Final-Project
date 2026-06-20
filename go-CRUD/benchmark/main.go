package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

const (
	defaultURL = "http://localhost:8080/students"
	defaultRequests = 10000
	defaultConcurrency = 50
)

func main() {
	url := envOr("BENCH_URL", defaultURL)
	total := envIntOr("BENCH_REQUESTS", defaultRequests)
	concurrency := envIntOr("BENCH_CONCURRENCY", defaultConcurrency)

	fmt.Printf("Go benchmark client\n")
	fmt.Printf("  Target:      %s\n", url)
	fmt.Printf("  Requests:    %d\n", total)
	fmt.Printf("  Concurrency: %d\n\n", concurrency)

	if err := warmup(url); err != nil {
		fmt.Fprintf(os.Stderr, "Server not reachable at %s: %v\n", url, err)
		os.Exit(1)
	}

	latencies := make([]time.Duration, 0, total)
	var latMu sync.Mutex
	var okCount atomic.Int64
	var errCount atomic.Int64

	start := time.Now()
	jobs := make(chan struct{}, total)
	for i := 0; i < total; i++ {
		jobs <- struct{}{}
	}
	close(jobs)

	var wg sync.WaitGroup
	for w := 0; w < concurrency; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{Timeout: 30 * time.Second}
			for range jobs {
				reqStart := time.Now()
				resp, err := client.Get(url)
				elapsed := time.Since(reqStart)

				latMu.Lock()
				latencies = append(latencies, elapsed)
				latMu.Unlock()

				if err != nil {
					errCount.Add(1)
					continue
				}
				io.Copy(io.Discard, resp.Body)
				resp.Body.Close()
				if resp.StatusCode == http.StatusOK {
					okCount.Add(1)
				} else {
					errCount.Add(1)
				}
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)

	printResults("Go", url, total, concurrency, elapsed, latencies, okCount.Load(), errCount.Load())
}

func warmup(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	return nil
}

func printResults(lang, url string, total, concurrency int, elapsed time.Duration, latencies []time.Duration, ok, failed int64) {
	sort.Slice(latencies, func(i, j int) bool { return latencies[i] < latencies[j] })

	rps := float64(total) / elapsed.Seconds()
	avg := average(latencies)
	p50 := percentile(latencies, 50)
	p95 := percentile(latencies, 95)
	p99 := percentile(latencies, 99)

	fmt.Println("Results")
	fmt.Println("-------")
	fmt.Printf("Language:     %s\n", lang)
	fmt.Printf("URL:          %s\n", url)
	fmt.Printf("Duration:     %v\n", elapsed.Round(time.Millisecond))
	fmt.Printf("Requests/sec: %.2f\n", rps)
	fmt.Printf("Successful:   %d\n", ok)
	fmt.Printf("Failed:       %d\n", failed)
	fmt.Printf("Avg latency:  %v\n", avg.Round(time.Microsecond))
	fmt.Printf("p50 latency:  %v\n", p50.Round(time.Microsecond))
	fmt.Printf("p95 latency:  %v\n", p95.Round(time.Microsecond))
	fmt.Printf("p99 latency:  %v\n", p99.Round(time.Microsecond))
}

func average(d []time.Duration) time.Duration {
	if len(d) == 0 {
		return 0
	}
	var sum time.Duration
	for _, v := range d {
		sum += v
	}
	return sum / time.Duration(len(d))
}

func percentile(sorted []time.Duration, p int) time.Duration {
	if len(sorted) == 0 {
		return 0
	}
	idx := (p * len(sorted)) / 100
	if idx >= len(sorted) {
		idx = len(sorted) - 1
	}
	return sorted[idx]
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envIntOr(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		var n int
		if _, err := fmt.Sscanf(v, "%d", &n); err == nil && n > 0 {
			return n
		}
	}
	return fallback
}
