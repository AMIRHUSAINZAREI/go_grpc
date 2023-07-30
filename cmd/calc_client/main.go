package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AMIRHUSAINZAREI/go_grpc_sample/pkg"
	pb "github.com/AMIRHUSAINZAREI/go_grpc_sample/proto/calc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Load port number from .env file
	port, err := pkg.GetEnv("CALC_GRPC_SERVER_PORT")
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cc := pb.NewCalcClient(conn)

	// Operation no
	var s int
	fmt.Println("1.Add\n2.Sub\n3.Mul\n4.Div")
	fmt.Scan(&s)

	// Operands
	var a, b int32
	fmt.Scan(&a)
	fmt.Scan(&b)

	switch s {
	case 1:
		res, err := cc.Add(ctx, &pb.Request{A: a, B: b})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.GetResult())
	case 2:
		res, err := cc.Sub(ctx, &pb.Request{A: a, B: b})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.GetResult())
	case 3:
		res, err := cc.Mul(ctx, &pb.Request{A: a, B: b})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.GetResult())
	case 4:
		res, err := cc.Div(ctx, &pb.Request{A: a, B: b})
		if err != nil {
			log.Fatal(err)
		}
		if res.GetError() != "" {
			log.Fatal(res.GetError())
		} else {
			fmt.Println(res.GetResult())
		}
	default:
		fmt.Println("operation number is between 1 and 4")
	}
}
