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

// PacketUserAdd inserts new user data.
type PacketUserAdd struct {
	UserID int
	Name   string
}

// PacketReviewAdd containes review data.
type PacketReviewAdd struct {
	UserID  int
	Score   int
	Comment string
}

// PacketRestaurantAdd contains data for new restaurant.
type PacketRestaurantAdd struct {
	Name      string
	Latitude  float64
	Longitude float64
	Altitude  float64
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
func (p *PacketReviewAdd) Type() PacketType {
	return PTReviewAdd
}

// Type implements Packet interface.
func (p *PacketRestaurantAdd) Type() PacketType {
	return PTRestaurantAdd
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
