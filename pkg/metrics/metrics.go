package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Определяем метрику для активных соединений
	ActiveConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "active_connections",
		Help: "Current number of active connections",
	})
	// Определяем метрику для общего числа запросов
	RequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "requests_total",
		Help: "Total number of requests",
	})
)

// Регистрируем метрики в Prometheus
func RegisterMetrics() {
	prometheus.MustRegister(ActiveConnections)
	prometheus.MustRegister(RequestsTotal)
}
