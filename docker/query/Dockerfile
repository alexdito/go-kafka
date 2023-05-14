FROM golang:1.20

ADD . /query

WORKDIR /query

RUN go mod tidy

RUN go mod download

RUN go build -o ./bin/query ./cmd/query/main.go

EXPOSE 8081

CMD [ "./bin/query" ]