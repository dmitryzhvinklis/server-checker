.PHONY: build run test

build:
	go build -o bin/custom_exporter ./cmd/custom_exporter

run:
	./bin/custom_exporter

test:
	go test ./...

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

start-prometheus:
	./scripts/start_prometheus.sh

start-grafana:
	./scripts/start_grafana.sh

start-alertmanager:
	./scripts/start_alertmanager.sh
