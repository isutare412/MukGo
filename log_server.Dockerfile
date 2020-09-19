FROM golang:1.15.2

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
COPY server server
COPY config/default/mukgo_log.yml /config/mukgo_log.yml

# Build api server binary
RUN go build -o log_server server/cmd/mukgo-log/mukgo-log.go

EXPOSE 7777

ENTRYPOINT [ "./log_server", "/config/mukgo_log.yml" ]
