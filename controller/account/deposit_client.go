package account

import (
	"context"
	pb "grpc-client/proto/account"
	"log"
	"time"

	"google.golang.org/grpc"
)

type DepositClient struct {
	conn *grpc.ClientConn
}

func NewDepositClient(conn *grpc.ClientConn) DepositClient {
	return DepositClient{
		conn: conn,
	}
}

func (d *DepositClient) Deposit(ctx context.Context, amount float32) (bool, error) {
	c := pb.NewDepositServiceClient(d.conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := c.Deposit(ctx, &pb.DepositRequest{Amount: amount})
	if err != nil {
		return r.GetOk(), err
	}

	return r.GetOk(), nil
}

func (d *DepositClient) GetDeposit(ctx context.Context) (float32, error) {
	c := pb.NewDepositServiceClient(d.conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := c.GetDeposit(ctx, &pb.GetDepositRequest{})
	if err != nil {
		log.Fatalf("Could not send deposit:%v", err)
	}

	return r.GetTotalDeposit(), nil
}
