package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/AlexKoppara/CarSim/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 8000, "server port")
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
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCarSimServer(grpcServer, &carSimServer{})
	grpcServer.Serve(lis)
}
