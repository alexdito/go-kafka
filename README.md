# go-kafka
Тестовый проект для изучения возможностей связки kafka + golang

1. Consumer cmd/consumer/main.go
2. Producer cmd/producer/main.go

Перед запуском producer и consumer нужно создать топик с помощью команд:
Зайти в консоль контейнера: 
`docker exec -it kafka /bin/sh`

Создать топик: 
`kafka-console-producer --broker-list 0.0.0.0:9092 --topic mytopic`

Слушать топик из контейнера kafka: 
`kafka-console-consumer --bootstrap-server 0.0.0.0:9092 --topic mytopic --from-beginning`

Для наглядности работы нужно запустить сначала zookeeper и kafka, а затем go-producer и go-consumer в разных консолях без флага -d (Не в режиме демона)