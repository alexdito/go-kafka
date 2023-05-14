FROM golang:1.20

ADD . /command

WORKDIR /command

RUN go mod tidy

RUN go mod download

RUN go build -o ./bin/command ./cmd/command/main.go

EXPOSE 8080

CMD [ "./bin/command" ]