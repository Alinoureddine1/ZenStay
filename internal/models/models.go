package models

import "time"

type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	RoomID    int
	StartDate string
	EndDate   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Processed int
}

type Users struct {
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

type Rooms struct {
	Id        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Restrictions struct {
	Id              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type Reservations struct {
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
	Room      Rooms
}

type RoomRestrictions struct {
	Id            int
	RoomId        int
	ReservationId int
	RestrictionId int
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time

	Room        Rooms
	Reservation Reservations
	Restriction Restrictions
}
