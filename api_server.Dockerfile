FROM golang:1.15.2

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
COPY server server
COPY config/mukgo_api.yml mukgo_api.yml

# Build api server binary
RUN go build -o api_server server/cmd/mukgo-api/mukgo-api.go

EXPOSE 7777

ENTRYPOINT [ "./api_server", "mukgo_api.yml" ]
