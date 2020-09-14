FROM golang:1.15.2

WORKDIR /workspace

COPY go.mod go.mod
COPY server server

# Build api server binary
RUN go build -o api-server server/cmd/mukgo-api/mukgo-api.go

EXPOSE 7777

ENTRYPOINT [ "./api-server" ]
