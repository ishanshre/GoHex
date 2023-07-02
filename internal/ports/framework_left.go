package ports

import (
	"context"

	"github.com/ishanshre/GoHex/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, res *pb.OperationParameters) (*pb.Answer, error)
	GetSubstraction(ctx context.Context, res *pb.OperationParameters) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, res *pb.OperationParameters) (*pb.Answer, error)
	GetDivision(ctx context.Context, res *pb.OperationParameters) (*pb.Answer, error)
}
