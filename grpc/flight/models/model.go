package models

import "time"

type GetListFlightRequest struct {
	Code string `json:"Code"`
}

type UpdateFlightRequest struct {
	From          string
	Destination   string
	Code          string
	TotalSlot     int
	DepartureTime time.Time
	ArrivalTime   time.Time
}
