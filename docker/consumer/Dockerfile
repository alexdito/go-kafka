FROM golang:1.19

ADD . /consumer

WORKDIR /consumer

RUN go mod download

RUN go build -o ./bin/consumer ./cmd/consumer/main.go

EXPOSE 8080
EXPOSE 80

CMD [ "./bin/consumer" ]