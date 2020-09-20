package server

import (
	"time"
)

// PacketType matches int with packet structure.
type PacketType int32

// All PacketTypes.
const (
	PTInvalid PacketType = iota
	PTReview
	PTLog
	PTAck
	PTError
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
	UserID  int    `json:"userid"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

// PacketAck contains ack response.
type PacketAck struct {
}

// PacketError contains error messge.
type PacketError struct {
	Message string
}

/******************************************************************************
* Log packets
******************************************************************************/

// PacketLog defines struct for log message among servers.
type PacketLog struct {
	Timestamp time.Time
	Msg       string
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
func (p *PacketError) Type() PacketType {
	return PTError
}

// Type implements Packet interface.
func (p *PacketLog) Type() PacketType {
	return PTLog
}
