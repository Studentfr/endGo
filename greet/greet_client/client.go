package greet_client

import (
	"context"
	"fmt"
	"goMod/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewCalculatorServiceClient(conn)
	doManyTimesFromServer(c)

}
func doManyTimesFromServer(c greetpb.GreetServiceClient) {
	ctx := context.Background()
	req := &greetpb.RequestNum{Number: 120}

	stream, err := c.Calculate(ctx, req)
	if err != nil {
		log.Fatalf("error while calling calculate RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we've reached the end of the stream
				break LOOP
			}
			log.Fatalf("error while reciving from calculate RPC %v", err)
		}
		log.Printf("response from calculate:%v \n", res.GetNumber())
	}

}
