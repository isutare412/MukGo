FROM golang:1.15.2

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
COPY server server

# Build api server binary
RUN go build -o log_server server/cmd/mukgo-log/mukgo-log.go

EXPOSE 7777

ENTRYPOINT [ "./log_server" ]
