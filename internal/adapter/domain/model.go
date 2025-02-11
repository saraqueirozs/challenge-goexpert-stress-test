package domain

// "result" será o resultado de uma requisição
type Result struct {
	StatusCode int
}

// "report" representa o relatório final do teste de carga
type Report struct {
	TotalTime  string      `json:"total_time"`
	TotalSend  int         `json:"total_send"`
	StatusCode map[int]int `json:"status_code"`
}
