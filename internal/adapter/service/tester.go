package service

import (
	"challenge-goexpert-stress-test/internal/adapter/domain"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LoadTester struct{}

func NewLoadTester() *LoadTester {
	return &LoadTester{}
}

func (lt *LoadTester) RunLoadTest(url string, totalRequests, concurrency int) domain.Report {
	var wg sync.WaitGroup
	results := make(chan domain.Result, totalRequests)

	startTime := time.Now()

	requestsPerWorker := totalRequests / concurrency
	extraRequests := totalRequests % concurrency

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		reqs := requestsPerWorker
		if i < extraRequests {
			reqs++
		}
		go lt.worker(url, reqs, results, &wg)
	}

	wg.Wait()
	close(results)

	return lt.generateReport(results, startTime)
}

func (lt *LoadTester) worker(url string, requests int, results chan<- domain.Result, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}

	for i := 0; i < requests; i++ {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Erro:", err)
			results <- domain.Result{StatusCode: 0}
			continue
		}
		results <- domain.Result{StatusCode: resp.StatusCode}
		resp.Body.Close()
	}
}

func (lt *LoadTester) generateReport(results <-chan domain.Result, startTime time.Time) domain.Report {
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	statusCount := make(map[int]int)
	totalSent := 0

	for result := range results {
		statusCount[result.StatusCode]++
		totalSent++
	}

	// Exibir distribuição dos status HTTP
	fmt.Println("\n=== Relatório de Teste de Carga ===")
	fmt.Printf("Tempo total gasto: %s\n", elapsedTime)
	fmt.Printf("Total de requisições enviadas: %d\n", totalSent)
	for status, count := range statusCount {
		fmt.Printf("Código HTTP %d: %d requisições\n", status, count)
	}

	return domain.Report{
		TotalTime:  elapsedTime.String(),
		TotalSend:  totalSent,
		StatusCode: statusCount,
	}
}
