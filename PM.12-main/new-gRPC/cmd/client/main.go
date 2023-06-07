package main

import (
	"context"
	"flag"
	"log"
	"new-gRPC/pkg/api"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 3 {
		log.Fatal(("Not enough arguments"))
	}

	name := flag.Arg((0))

	x, err := strconv.Atoi(flag.Arg((1)))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg((2)))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewAdderClient(conn)

	res, err := c.NewConnect(context.Background(), &api.NewConnRequest{Name: name})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("You connected as %v", res.GetResult())

	ress, err := c.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(ress.GetResult())
}
