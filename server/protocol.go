package server

// RabbitMQ exchange names.
const (
	// MGLogs defines exchange name to log server.
	MGLogs string = "mglogs"

	// MGDB defines exchange between api server and db server.
	MGDB string = "mgdb"
)

// RabbitMQ queue names.
const (
	Log    string = "log"
	API2DB string = "api2db"
	DB2API string = "db2api"
)

// Keys used in amqp.Table.
const (
	MsgType string = "MsgType"
	Sender  string = "Sender"
)

// Possible senders of amqp.Table.
const (
	API string = "api"
	DB  string = "db"
)
