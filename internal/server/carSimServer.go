package server

import (
	"context"
	"fmt"

	pb "github.com/AlexKoppara/CarSim/api"
)

type carSimServer struct {
	pb.UnimplementedCarSimServer
}

// CreateCar creates a new in memory car
func (s *carSimServer) CreateCar(ctx context.Context, carInfo *pb.CarInfo) (*pb.Point, error) {
	fmt.Println(carInfo.GetName())

	return &pb.Point{Lat: 1, Long: 1}, nil
}

func (s *carSimServer) DriveCar(accVector *pb.AccVector, stream pb.CarSim_DriveCarServer) error {

}
