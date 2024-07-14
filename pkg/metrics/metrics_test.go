package metrics

import (
	"testing"
)

func TestRegisterMetrics(t *testing.T) {
	// Проверяем, что функция RegisterMetrics не вызывает ошибок
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterMetrics() panicked: %v", r)
		}
	}()
	RegisterMetrics()
}
