package server

import (
	"time"
)

// PacketType matches int with packet structure.
type PacketType int

// All PacketTypes.
const (
	PTReview PacketType = iota
	PTLog
	PTAck
)

// Packet interface can generate unique PacketType.
type Packet interface {
	Type() PacketType
}

/******************************************************************************
* Database packets
******************************************************************************/

// PacketReview containes review data.
type PacketReview struct {
	UserID  int
	Score   int
	Comment string
}

// PacketAck contains ack response.
type PacketAck struct {
}

/******************************************************************************
* Log packets
******************************************************************************/

// PacketLog defines struct for log message among servers.
type PacketLog struct {
	Timestamp time.Time `json:"timestamp"`
	Msg       string    `json:"message"`
}

/******************************************************************************
* Packet interface methods
******************************************************************************/

// Type implements Packet interface.
func (p *PacketReview) Type() PacketType {
	return PTReview
}

// Type implements Packet interface.
func (p *PacketAck) Type() PacketType {
	return PTAck
}

// Type implements Packet interface.
func (p *PacketLog) Type() PacketType {
	return PTLog
}
