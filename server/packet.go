package server

import (
	"time"

	"github.com/isutare412/MukGo/server/common"
	"github.com/isutare412/MukGo/server/console"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PacketType matches int with packet structure.
type PacketType int32

// ErrorType enumerates error types in error packet.
type ErrorType int32

// All PacketTypes.
const (
	// api server to dabase server
	PTInvalid PacketType = iota
	PTADUserAdd
	PTADUserGet
	PTADReviewGet
	PTADReviewsGet
	PTADReviewAdd
	PTADReviewDel
	PTADRestaurantGet
	PTADRestaurantAdd
	PTADRestaurantsGet
	PTADRestaurantsAdd
	PTADRankingGet
	PTADLikeAdd
	PTADLikeDel

	// database server to api server
	PTDAAck
	PTDAError
	PTDAUser
	PTDAUsers
	PTDARestaurant
	PTDARestaurants
	PTDAReviews
	PTDAReview

	// log packet type
	PTLog
)

// All ErrorTypes.
const (
	ETInvalid ErrorType = iota
	ETInternal
	ETUserExists
	ETNoSuchUser
	ETNoSuchRestaurant
	ETLikeExists
	ETNoPermission
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
	UserID       string
	HeavyRequest bool
}

// ADPacketReviewGet containes review data.
type ADPacketReviewGet struct {
	UserID   string
	ReviewID primitive.ObjectID
}

// ADPacketReviewsGet containes review data.
type ADPacketReviewsGet struct {
	UserID string
	RestID primitive.ObjectID
}

// ADPacketReviewAdd containes review data.
type ADPacketReviewAdd struct {
	UserID    string
	RestID    primitive.ObjectID
	Score     int32
	Comment   string
	Menus     []string
	Wait      bool
	NumPeople int32
	Timestamp int64
}

// ADPacketReviewDel containes review data.
type ADPacketReviewDel struct {
	UserID   string
	ReviewID primitive.ObjectID
}

// ADPacketRestaurantGet request restaurant data.
type ADPacketRestaurantGet struct {
	RestID primitive.ObjectID
}

// ADPacketRestaurantAdd contains data for new restaurant.
type ADPacketRestaurantAdd struct {
	*common.Restaurant
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

// ADPacketRankingGet requests user ranking.
type ADPacketRankingGet struct {
}

// ADPacketLikeAdd requests like on specific review.
type ADPacketLikeAdd struct {
	UserID   string
	ReviewID primitive.ObjectID
}

// ADPacketLikeDel requests cancel of like on specific review.
type ADPacketLikeDel struct {
	UserID   string
	ReviewID primitive.ObjectID
}

/******************************************************************************
* Database to API packets
******************************************************************************/

// DAPacketAck contains ack response.
type DAPacketAck struct {
}

// DAPacketError contains error messge.
type DAPacketError struct {
	ErrorType
}

// DAPacketUser contains user data.
type DAPacketUser struct {
	*common.User
}

// DAPacketUsers contains users ordered by certain criterion.
type DAPacketUsers struct {
	Users []*common.User
}

// DAPacketRestaurant contains multiple Restaurant models.
type DAPacketRestaurant struct {
	Restaurant *common.Restaurant
}

// DAPacketRestaurants contains multiple Restaurant models.
type DAPacketRestaurants struct {
	Restaurants []*common.Restaurant
}

// DAPacketReviews contains multiple Review models.
type DAPacketReviews struct {
	Reviews []*common.Review
}

// DAPacketReview contains review model.
type DAPacketReview struct {
	*common.Review
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
func (p *ADPacketReviewGet) Type() PacketType {
	return PTADReviewGet
}

// Type implements Packet interface.
func (p *ADPacketReviewsGet) Type() PacketType {
	return PTADReviewsGet
}

// Type implements Packet interface.
func (p *ADPacketReviewAdd) Type() PacketType {
	return PTADReviewAdd
}

// Type implements Packet interface.
func (p *ADPacketReviewDel) Type() PacketType {
	return PTADReviewDel
}

// Type implements Packet interface.
func (p *ADPacketRestaurantGet) Type() PacketType {
	return PTADRestaurantGet
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
func (p *ADPacketRankingGet) Type() PacketType {
	return PTADRankingGet
}

// Type implements Packet interface.
func (p *ADPacketLikeAdd) Type() PacketType {
	return PTADLikeAdd
}

// Type implements Packet interface.
func (p *ADPacketLikeDel) Type() PacketType {
	return PTADLikeDel
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
func (p *DAPacketUser) Type() PacketType {
	return PTDAUser
}

// Type implements Packet interface.
func (p *DAPacketUsers) Type() PacketType {
	return PTDAUsers
}

// Type implements Packet interface.
func (p *DAPacketRestaurant) Type() PacketType {
	return PTDARestaurant
}

// Type implements Packet interface.
func (p *DAPacketRestaurants) Type() PacketType {
	return PTDARestaurants
}

// Type implements Packet interface.
func (p *DAPacketReviews) Type() PacketType {
	return PTDAReviews
}

// Type implements Packet interface.
func (p *DAPacketReview) Type() PacketType {
	return PTDAReview
}

// Type implements Packet interface.
func (p *PacketLog) Type() PacketType {
	return PTLog
}

/******************************************************************************
* ErrorType string interface
******************************************************************************/

func (e ErrorType) String() string {
	var msg string
	switch e {
	case ETInternal:
		msg = "internal error"
	case ETUserExists:
		msg = "user exists"
	case ETNoSuchUser:
		msg = "no such user"
	case ETNoSuchRestaurant:
		msg = "no such restaurant"
	default:
		msg = "unknown error"
	}
	return msg
}
