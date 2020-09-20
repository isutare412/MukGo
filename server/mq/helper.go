package mq

import (
	"fmt"

	"github.com/isutare412/MukGo/server"
	"github.com/streadway/amqp"
)

// ParseHeader parses custom data from RabbitMQ delivery header.
func ParseHeader(h amqp.Table) (
	sender string, packetType server.PacketType, err error,
) {
	if h == nil {
		err = fmt.Errorf("header does not exists in delievery")
		return
	}

	msgType, ok := h[server.MsgType].(int32)
	if ok {
		packetType = server.PacketType(msgType)
	}

	name, ok := h[server.Sender].(string)
	if ok {
		sender = name
	}

	return
}
