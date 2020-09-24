FROM golang:1.15.2

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
COPY server server

# Build database server binary
RUN go build -o db_server server/cmd/mukgo-db/mukgo-db.go

EXPOSE 7777

ENTRYPOINT [ "./db_server" ]
