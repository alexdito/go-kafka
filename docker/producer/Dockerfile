FROM golang:1.20

ADD . /producer

WORKDIR /producer

RUN go mod download

RUN go build -o ./bin/producer ./cmd/producer/main.go

EXPOSE 8080
EXPOSE 80

CMD [ "./bin/producer" ]