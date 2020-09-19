package server

import "time"

// PacketLog defines struct for log message among servers.
type PacketLog struct {
	Timestamp time.Time `json:"timestamp"`
	Msg       string    `json:"message"`
}

// RabbitMQ exchange names.
const (
	// MGLogs defines exchange name to log server.
	MGLogs string = "mglogs"

	// MGDB defines exchange between api server and db server.
	MGDB string = "mgdb"
)

// RabbitMQ queue names.
const (
	Log     string = "log"
	APIToDB string = "api2db"
	DBToAPI string = "db2api"
)

// Keys used in amqp.Table.
const (
	Sender string = "sender"
)

// Possible senders of amqp.Table.
const (
	API string = "api"
)
