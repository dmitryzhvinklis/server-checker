 # Система мониторинга серверов

Этот проект представляет собой систему мониторинга серверов, включающую кастомный экспортёр на Go, Prometheus, Grafana и Alertmanager.


## Как запустить проект

1. Соберите проект:

    ```bash
    make build
    ```

2. Запустите Docker Compose:

    ```bash
    make docker-up
    ```

3. Запустите кастомный экспортёр:

    ```bash
    make run
    ```

4. Проверьте работу сервисов:

    - Prometheus: [http://localhost:9090](http://localhost:9090)
    - Grafana: [http://localhost:3000](http://localhost:3000)
    - Alertmanager: [http://localhost:9093](http://localhost:9093)
    - Кастомный экспортёр: [http://localhost:8080](http://localhost:8080)

5. Проверьте метрики в Prometheus:

    Откройте [http://localhost:9090](http://localhost:9090) и введите `custom_exporter` в строке поиска метрик.

## Как тестировать проект

Запустите тесты:

```bash
make test
```