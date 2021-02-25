package greet_server

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	greetpb.UnimplementedCalculatorServiceServer
}

func (s *Server) Calculate(req *greetpb.RequestNum,
	stream greetpb.CalculatorService_CalculateServer) error {
	fmt.Printf("Calculate function was invoked with %v \n", req)
	number := req.GetX()
	for number > 1 {
		num := getFirstPrime(number)
		number /= num
		res := &greetpb.ResponseNum{Y: num}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending greet many times responses: %v", err.Error())
		}
		time.Sleep(time.Second)
	}
	return nil
}

func getFirstPrime(number int32) int32 {
	for i := 2; int32(i) <= number; i++ {
		if number%int32(i) == 0 {
			return int32(i)
		}
	}
	return number
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Server is running on port:8080")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
