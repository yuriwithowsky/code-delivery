# Simulador de corridas

## Como iniciar o simulador?

1. Inicie o kafka

```bash
cd .docker/kafka
docker-compose up -d
docker exec -it kafka_kafka_1 bash

```

2. Inicie o simulador

```bash
docker-compose up -d
docker exec -it simulator-service bash

go run main.go
```

## Bonus: Dentro do container do Kafka

```bash
# criar um consumer no console
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal

# criar um producer no console
kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
```

- Teste de mensagens

```json
{"clientId": "1", "routeId": "1"}
{"clientId": "2", "routeId": "2"}
{"clientId": "3", "routeId": "3"}
```
