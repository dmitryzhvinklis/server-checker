package main

import (
	"log"
	"net/http"

	"server-checker/pkg/handlers"
	"server-checker/pkg/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Регистрируем метрики
	metrics.RegisterMetrics()

	// Обработчик для метрик Prometheus
	http.Handle("/metrics", promhttp.Handler())

	// Обработчик для основных запросов
	http.HandleFunc("/", handlers.RequestHandler)

	// Запуск HTTP сервера на порту 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
