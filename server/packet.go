package server

import (
	"time"

	"github.com/isutare412/MukGo/server/common"
	"github.com/isutare412/MukGo/server/console"
)

// PacketType matches int with packet structure.
type PacketType int32

// All PacketTypes.
const (
	PTInvalid PacketType = iota
	PTUserAdd
	PTReviewAdd
	PTRestaurantAdd
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

// ADPacketUserAdd inserts new user data.
type ADPacketUserAdd struct {
	UserID int
	Name   string
}

// ADPacketReviewAdd containes review data.
type ADPacketReviewAdd struct {
	UserID  int
	Score   int
	Comment string
}

// ADPacketRestaurantAdd contains data for new restaurant.
type ADPacketRestaurantAdd struct {
	Name  string
	Coord common.Coordinate
}

// DAPacketAck contains ack response.
type DAPacketAck struct {
}

// DAPacketError contains error messge.
type DAPacketError struct {
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
func (p *ADPacketUserAdd) Type() PacketType {
	return PTUserAdd
}

// Type implements Packet interface.
func (p *ADPacketReviewAdd) Type() PacketType {
	return PTReviewAdd
}

// Type implements Packet interface.
func (p *ADPacketRestaurantAdd) Type() PacketType {
	return PTRestaurantAdd
}

// Type implements Packet interface.
func (p *DAPacketAck) Type() PacketType {
	return PTAck
}

// Type implements Packet interface.
func (p *DAPacketError) Type() PacketType {
	return PTError
}

// Type implements Packet interface.
func (p *PacketLog) Type() PacketType {
	return PTLog
}
