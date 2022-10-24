package repositories

import (
	"flightbooking/db"

	"gorm.io/gorm"
)

type IFlightService interface {
	GetList() ([]db.Flight, error)
	Get(code string) (*db.Flight, error)
	GetById(id int) (*db.Flight, error)
}

type FlightService struct {
	DB *gorm.DB
}

func NewFlightService() FlightService {
	var db = db.NewGormDB()
	return FlightService{db}
}

func (f *FlightService) GetList() ([]db.Flight, error) {
	var query []db.Flight
	find := f.DB.Find(&query)
	if find.Error != nil {
		return nil, find.Error
	}
	return query, nil
}

func (f *FlightService) Get(code string) (*db.Flight, error) {
	query := db.Flight{}
	find := f.DB.Find(&query, db.Flight{Code: code})
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}

func (f *FlightService) GetById(id int) (*db.Flight, error) {
	query := db.Flight{}
	find := f.DB.Find(&query, db.Flight{ID: id})
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}
