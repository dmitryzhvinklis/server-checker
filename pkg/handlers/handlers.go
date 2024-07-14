package handlers

import (
	"net/http"

	"server-checker/pkg/metrics"
)

// Обработчик для основного запроса
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	// Увеличиваем счетчик запросов
	metrics.RequestsTotal.Inc()
	// Увеличиваем количество активных соединений
	metrics.ActiveConnections.Inc()
	defer metrics.ActiveConnections.Dec()
	// Отправляем ответ
	w.Write([]byte("Hello, World!"))
}
