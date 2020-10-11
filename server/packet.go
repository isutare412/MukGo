package server

import (
	"time"

	"github.com/isutare412/MukGo/server/console"
)

// PacketType matches int with packet structure.
type PacketType int32

// All PacketTypes.
const (
	PTInvalid PacketType = iota
	PTUserAdd
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

// PacketUserAdd inserts new user data.
type PacketUserAdd struct {
	UserID int
	Name   string
}

// PacketReview containes review data.
type PacketReview struct {
	UserID  int
	Score   int
	Comment string
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
	LogLevel  console.Level
	Msg       string
}

/******************************************************************************
* Packet interface methods
******************************************************************************/

// Type implements Packet interface.
func (p *PacketUserAdd) Type() PacketType {
	return PTUserAdd
}

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
