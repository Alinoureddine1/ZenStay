package models

type Reservation struct {
	// ID is the unique id of the reservation
	ID int
	// FirstName is the first name of the person making the reservation
	FirstName string
	// LastName is the last name of the person making the reservation
	LastName string
	// Email is the email of the person making the reservation
	Email string
	// Phone is the phone number of the person making the reservation
	Phone string
	// RoomID is the id of the room reserved
	RoomID int
	// Room is the room info of the room reserved
	//Room Room
	// StartDate is the start date of the reservation
	StartDate string
	// EndDate is the end date of the reservation
	EndDate string
	// CreatedAt is the time the reservation was created
	CreatedAt string
	// UpdatedAt is the time the reservation was updated
	UpdatedAt string
	// DeletedAt is the time the reservation was deleted
	DeletedAt string
	// Processed is true if the reservation has been processed
	Processed int
}
