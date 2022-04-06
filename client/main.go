package main

import (
	"context"
	cont "grpc-client/controller/account"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	//Check Args
	if len(os.Args) <= 1 {
		log.Println("Please enter Args")
		log.Fatal("Use get to retrieve account balance, and send to send amount")
		return
	}

	if os.Args[1] != "get" && os.Args[1] != "send" {
		log.Println("Invalid Args")
		log.Fatal("Use get to retrieve account balance, and send to send amount")
	}

	if os.Args[1] == "send" && len(os.Args) <= 2 {
		log.Fatal("Please enter the amount")
		return
	}

	if os.Args[1] == "send" && len(os.Args) >= 3 {
		_, err := strconv.ParseFloat(os.Args[2], 32)
		if err != nil {
			log.Fatal("Invalid Amount Args, Must be number")
		}
	}
	//Establish Conn
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect :%v", err)
	}
	defer conn.Close()

	d := cont.NewDepositClient(conn)

	if os.Args[1] == "send" {
		amount, _ := strconv.ParseFloat(os.Args[2], 32)
		res, err := d.Deposit(context.Background(), float32(amount))
		if err != nil {
			log.Fatalf("Failed to send: %v", err)
		}
		if res == true {
			log.Println("Balance Added")
			return
		}

	}

	res, err := d.GetDeposit(context.Background())
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	log.Printf("Current Balance: %v", res)
}
