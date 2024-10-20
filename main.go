package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type InvestmentData struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

func StartInvestmentServer() {
	http.HandleFunc("/api/v1/investment", func(w http.ResponseWriter, r *http.Request) {
		delay := rand.Intn(1000)
		time.Sleep(time.Duration(delay) * time.Millisecond)

		// Simula falhas ocasionais (por exemplo, 1 em cada 4 tentativas falha)
		if rand.Float32() < 0.25 {
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
			return
		}

		data := InvestmentData{
			Ticker: "AAPL",
			Price:  rand.Float64()*100 + 100,
		}
		response, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	fmt.Println("Investment Server running on port 8081")
	http.ListenAndServe(":8081", nil)
}

func main() {
	StartInvestmentServer()
}
