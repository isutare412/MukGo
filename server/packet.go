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
	PTADUserAdd
	PTADUserGet
	PTADReviewAdd
	PTADRestaurantAdd
	PTDAAck
	PTDAError
	PTDAUserExist
	PTDANoSuchUser
	PTDAUser
	PTLog
)

// Packet interface can generate unique PacketType.
type Packet interface {
	Type() PacketType
}

/******************************************************************************
* API to Database packets
******************************************************************************/

// ADPacketUserAdd inserts new user data.
type ADPacketUserAdd struct {
	UserID string
	Name   string
}

// ADPacketUserGet request user data.
type ADPacketUserGet struct {
	UserID string
}

// ADPacketReviewAdd containes review data.
type ADPacketReviewAdd struct {
	UserID  string
	Score   int
	Comment string
}

// ADPacketRestaurantAdd contains data for new restaurant.
type ADPacketRestaurantAdd struct {
	Name  string
	Coord common.Coordinate
}

/******************************************************************************
* Database to API packets
******************************************************************************/

// DAPacketAck contains ack response.
type DAPacketAck struct {
}

// DAPacketError contains error messge.
type DAPacketError struct {
	Message string
}

type DAPacketUserExist struct {
	UserID string
}

// DAPacketNoSuchUser contains error messge.
type DAPacketNoSuchUser struct {
	UserID string
}

// DAPacketUser contains error messge.
type DAPacketUser struct {
	UserID string
	Name   string
	Exp    int64
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
	return PTADUserAdd
}

// Type implements Packet interface.
func (p *ADPacketUserGet) Type() PacketType {
	return PTADUserGet
}

// Type implements Packet interface.
func (p *ADPacketReviewAdd) Type() PacketType {
	return PTADReviewAdd
}

// Type implements Packet interface.
func (p *ADPacketRestaurantAdd) Type() PacketType {
	return PTADRestaurantAdd
}

// Type implements Packet interface.
func (p *DAPacketAck) Type() PacketType {
	return PTDAAck
}

// Type implements Packet interface.
func (p *DAPacketError) Type() PacketType {
	return PTDAError
}

// Type implements Packet interface.
func (p *DAPacketUserExist) Type() PacketType {
	return PTDAUserExist
}

// Type implements Packet interface.
func (p *DAPacketNoSuchUser) Type() PacketType {
	return PTDANoSuchUser
}

// Type implements Packet interface.
func (p *DAPacketUser) Type() PacketType {
	return PTDAUser
}

// Type implements Packet interface.
func (p *PacketLog) Type() PacketType {
	return PTLog
}
