package repositories

import (
	"flightbooking/db"
	"log"

	"gorm.io/gorm"
)

type IBookingService interface {
	GetList(username string) ([]db.Booking, error)
	GetById(id int) (*db.Booking, error)
	Create(booking *db.Booking) (*db.Booking, error)
	GetByFlightId(flightId int) ([]db.Booking, error)
}

type BookingService struct {
	DB *gorm.DB
}

func NewBookingService() BookingService {
	var db = db.NewGormDB()
	return BookingService{DB: db}
}

func (b *BookingService) GetList(username string) ([]db.Booking, error) {
	var booking []db.Booking
	find := b.DB.Find(&booking, db.Booking{CustomerFullName: username})
	if find.Error != nil {
		return nil, find.Error
	}
	return booking, nil
}

func (b *BookingService) GetById(id int) (*db.Booking, error) {
	booking := db.Booking{}
	find := b.DB.Find(&booking, db.Booking{ID: id})
	if find.Error != nil {
		return nil, find.Error
	}
	return &booking, nil
}

func (b *BookingService) Create(booking *db.Booking) (*db.Booking, error) {
	create := b.DB.Create(booking)
	if create.Error != nil {
		return nil, create.Error
	}
	return booking, nil
}

func (b *BookingService) GetByFlightId(flightId int) ([]db.Booking, error) {
	var result []db.Booking
	find := b.DB.Find(&result, db.Booking{FlightId: flightId})
	if find.Error != nil {
		log.Println("Error occurred in GetByFlightId", find.Error)
		return nil, find.Error
	}
	return result, nil
}
