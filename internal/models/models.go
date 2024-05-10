package models

import "time"


type Reservation struct {
	FirstName string
	LastName string
	Email string
	Phone string
}

// Users is the model for the user table in the database
type Users struct {
	ID int
	FirstName string
	LastName string
	Email string
	Password string
	AccessLevel int
	CreatedAt time.Time
	UpdatedAt time.Time
}
// Rooms is the model for the rooms table in the database
type Rooms struct {
	ID int
	RoomName string
	CreatedAt time.Time
	UpdatedAt time.Time
}
// Restrictions is the model for the restrictions table in the database
type Restrictions struct {
	ID int
	RestrictionName string
	CreatedAt time.Time
	UpdatedAt time.Time
}
// Reservations is the model for the reservations table in the database
type Reservations struct {
	ID int
	FirstName string
	LastName string
	Email string
	Phone string
	StartDate time.Time
	EndDate time.Time
	RoomID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room Rooms
}
// RoomRestrictions is the model for the room_restrictions table in the database
type RoomRestrictions struct {
	ID int
	StartDate time.Time
	EndDate time.Time
	RoomID int
	ReservationID int
	RestrictionID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room Rooms
	Reservation Reservations
	Restriction Restrictions
}
