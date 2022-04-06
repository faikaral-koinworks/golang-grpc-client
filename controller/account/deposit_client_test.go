package account_test

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"

	cont "grpc-client/controller/account"
	pb "grpc-client/proto/account"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type mockDepositServiceServer struct {
	pb.UnimplementedDepositServiceServer
}

var accountValue float32 = 0

func (d *mockDepositServiceServer) Deposit(c context.Context, in *pb.DepositRequest) (*pb.DepositResponse, error) {
	if in.GetAmount() <= 0 {
		return &pb.DepositResponse{Ok: false}, errors.New("cannot deposit negative balance")
	}
	accountValue = accountValue + in.GetAmount()
	return &pb.DepositResponse{Ok: true}, nil
}

func (d *mockDepositServiceServer) GetDeposit(c context.Context, in *pb.GetDepositRequest) (*pb.GetDepositResponse, error) {
	return &pb.GetDepositResponse{TotalDeposit: accountValue}, nil
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterDepositServiceServer(server, &mockDepositServiceServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}

}

func TestDepositServiceClient_GetDeposit(t *testing.T) {
	test := struct {
		name string
		res  float32
		err  string
	}{
		"Valid Test Get Deposit",
		0,
		"",
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	t.Run(test.name, func(t *testing.T) {
		dclient := cont.NewDepositClient(conn)
		response, err := dclient.GetDeposit(context.Background())

		if response != test.res {
			t.Error("error : expected", test.res, "received:", response)
		}
		if err != nil {
			if er, _ := status.FromError(err); er.Message() != test.err {
				t.Error("error msg expected: ", test.err, "received", er.Message())
			}
		}
	})
}

func TestDepositServiceClient_Deposit(t *testing.T) {
	test := []struct {
		name   string
		amount float32
		res    bool
		err    string
	}{
		{
			"Invalid Request With Negative Amount",
			-1,
			false,
			"cannot deposit negative balance",
		},
		{
			"Valid Request With Valid Amount",
			1,
			true,
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			dclient := cont.NewDepositClient(conn)
			response, err := dclient.Deposit(context.Background(), v.amount)

			if response != v.res {
				t.Error("error : expected", v.res, "received:", response)
			}
			if err != nil {
				if er, _ := status.FromError(err); er.Message() != v.err {
					t.Error("error msg expected: ", v.err, "received", er.Message())
				}
			}
		})
	}
}
