package handlers

import (
	"context"
	"flightbooking/grpc/flight/repositories"
	"flightbooking/proto"
)

type Server struct {
	proto.FlightServiceServer
	FlightService repositories.IFlightService
}

func (s *Server) SearchFlight(ctx context.Context, req *proto.QueryFlightDetail) (*proto.FlightDetailResponse, error) {
	found, err := s.FlightService.GetById(int(req.FlightId))
	if err != nil {
		return nil, err
	}
	return &proto.FlightDetailResponse{
		Id:            int32(found.ID),
		Code:          found.Code,
		TotalSlot:     int32(found.TotalSlot),
		DepartureTime: found.DepartureTime.String(),
		ArrivalTime:   found.ArrivalTime.String(),
	}, nil
}
