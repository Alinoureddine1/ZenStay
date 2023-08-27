package models

import "time"

type User struct {
	Id          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Room struct {
	Id        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Restriction struct {
	Id              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type Reservation struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	RoomId    int
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Room      Room
}

type RoomRestriction struct {
	Id            int
	RoomId        int
	ReservationId int
	RestrictionId int
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time

	Room        Room
	Reservation Reservation
	Restriction Restriction
}
