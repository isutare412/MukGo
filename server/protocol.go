package server

// RabbitMQ exchange names.
const (
	// MGLogs defines exchange name to log server.
	MGLogs string = "mglogs"
)

// RabbitMQ queue names.
const (
	Log string = "log"
)

// Keys used in amqp.Table.
const (
	Sender string = "sender"
)

// Possible senders of amqp.Table.
const (
	API string = "api"
)

// PacketLog defines struct for log message among servers.
type PacketLog struct {
	Msg string `json:"message"`
}
