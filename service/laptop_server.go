package service

import (
	"context"
	"errors"
	"grpc/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	laptopStore LaptopStore
}

func NewLaptopServer(laptopStore LaptopStore) *LaptopServer {
	return &LaptopServer{laptopStore}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("Receive create laptop request with id: %s", laptop.Id)
	if len(laptop.Id) > 0 {
		// cek apakah valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID tidak valid: %v", err)
		}
	} else {
		// jika id laptop kosong, generate uuid baru
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "gagal generate ID: %v", err)
		}
		laptop.Id = id.String()
	}
	if ctx.Err() == context.Canceled {
		log.Print("request dibatalkan")
		return nil, status.Error(codes.Canceled, "request dibatalkan")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
	}

	err := server.laptopStore.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "gagal menyimpan laptop ke store: %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

func (server *LaptopServer) Hello(ctx context.Context, req *pb.CreateHelloRequest) (*pb.CreateHelloResponse, error) {
	log.Printf("Received: %s", req.GetName())
	return &pb.CreateHelloResponse{Reply: "Hello " + req.GetName()}, nil
}

func (server *LaptopServer) Find(ctx context.Context, req *pb.CreateFindRequest) (*pb.CreateFindResponse, error) {
	log.Printf("Received: %s", req.GetId())
	laptop, err := server.laptopStore.Find(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "id tidak ditemukan")
	}

	return &pb.CreateFindResponse{Laptop: laptop}, nil
}
