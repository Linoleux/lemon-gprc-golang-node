package main

import (
	"context"
	"flag"
	"fmt"
	"grpc/pb"
	"grpc/sample"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	serverAddress = flag.String("address", "localhost:50051", "the server address")
	name          = flag.String("name", "world", "username")
)

func main() {
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	defer conn.Close()

	laptopClient := pb.NewLaptopServiceClient(conn)
	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}
	log.Printf("created laptop with id: %s", res.Id)

	x, _ := laptopClient.Hello(ctx, &pb.CreateHelloRequest{
		Name: *name,
	})
	fmt.Println(x)

	// res, err := laptopClient.Find(ctx, &pb.CreateFindRequest{
	// 	Id: "1cc352b0-568f-427d-b83d-4ba9a0354432",
	// })
	// if err != nil {
	// 	log.Print("laptop not found")
	// }

	// fmt.Println(res.Laptop)
}
