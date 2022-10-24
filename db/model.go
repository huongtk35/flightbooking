package db

import "time"

const BookingStatusCreated = "Created"

type Booking struct {
	ID               int       `gorm:"column=id;not null;primaryKey"`
	Username         string    `gorm:"column=userName;not null"`
	CustomerFullName string    `gorm:"column=customerFullName;not null"`
	FlightId         int       `gorm:"column=flightId;not null"`
	Code             string    `gorm:"column=code;not null"`
	Status           string    `gorm:"column=status;not null"`
	ReservedSlot     int       `gorm:"reservedSlot;not null"`
	BookedDate       time.Time `gorm:"column=bookedDate;not null"`
	CreatedAt        time.Time `gorm:"column=createdAt;not null"`
	UpdatedAt        time.Time `gorm:"column=updatedAt;not null"`
}

func (b *Booking) TableName() string {
	return "dev.booking"
}

type Flight struct {
	ID            int       `gorm:"column=id;not null;primaryKey"`
	From          string    `gorm:"column=from;not null"`
	Destination   string    `gorm:"column=from;not null"`
	Code          string    `gorm:"column=code;not null"`
	TotalSlot     int       `gorm:"column=total_slot;not null"`
	DepartureTime time.Time `gorm:"column=departure_time;not null"`
	ArrivalTime   time.Time `gorm:"column=arrival_time;not null"`
}

func (f *Flight) TableName() string {
	return "dev.flight"
}

type Gender int

const (
	Female Gender = iota
	Male
)

type User struct {
	Username  string    `gorm:"column=username;not null;primaryKey"`
	Password  string    `gorm:"column=password;not null"`
	FirstName string    `gorm:"column=first_name;not null"`
	LastName  string    `gorm:"column=last_name;not null"`
	Gender    Gender    `gorm:"column=gender;not null"`
	CreatedAt time.Time `gorm:"column=created_at;not null"`
	UpdatedAt time.Time `gorm:"column=updated_at;not null"`
}

func (u *User) TableName() string {
	return "dev.user"
}
