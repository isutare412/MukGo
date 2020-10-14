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
	// api server to dabase server
	PTInvalid PacketType = iota
	PTADUserAdd
	PTADUserGet
	PTADReviewAdd
	PTADRestaurantAdd
	PTADRestaurantsGet
	PTADRestaurantsAdd

	// database server to api server
	PTDAAck
	PTDAError
	PTDAUserExist
	PTDANoSuchUser
	PTDAUser
	PTDARestaurants
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

// ADPacketRestaurantsGet request restaurants within user's sight.
type ADPacketRestaurantsGet struct {
	UserID string
	Coord  common.Coordinate
}

// ADPacketRestaurantsAdd contains data for new restaurants.
type ADPacketRestaurantsAdd struct {
	Restaurants []*common.Restaurant
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

// DAPacketUserExist contains error message.
type DAPacketUserExist struct {
	UserID string
}

// DAPacketNoSuchUser contains error messge.
type DAPacketNoSuchUser struct {
	UserID string
}

// DAPacketUser contains user data.
type DAPacketUser struct {
	UserID string
	Name   string
	Exp    int64
}

// DAPacketRestaurants contains multiple Restaurant models.
type DAPacketRestaurants struct {
	Restaurants []*common.Restaurant
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
func (p *ADPacketRestaurantsGet) Type() PacketType {
	return PTADRestaurantsGet
}

// Type implements Packet interface.
func (p *ADPacketRestaurantsAdd) Type() PacketType {
	return PTADRestaurantsAdd
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
func (p *DAPacketRestaurants) Type() PacketType {
	return PTDARestaurants
}

// Type implements Packet interface.
func (p *PacketLog) Type() PacketType {
	return PTLog
}
