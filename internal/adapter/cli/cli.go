package cli

import (
	"flag"
	"fmt"

	"challenge-goexpert-stress-test/internal/adapter/service"
)

func RunCLI(tester *service.LoadTester) {
	url := flag.String("url", "", "URL do serviço a ser testado")
	totalRequests := flag.Int("requests", 100, "Número total de requisições")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("A URL deve ser informada com o parâmetro --url")
		return
	}

	report := tester.RunLoadTest(*url, *totalRequests, *concurrency)

	fmt.Println("\n=== Relatório de Teste de Carga ===")
	fmt.Printf("Tempo total gasto: %s\n", report.TotalTime)
	fmt.Printf("Total de requisições enviadas: %d\n", report.TotalSend)
	for status, count := range report.StatusCode {
		fmt.Printf("Código HTTP %d: %d requisições\n", status, count)
	}
}
