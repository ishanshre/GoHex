package ports

import (
	"context"
)

type GRPCPort interface {
	Run()
	GetAddition()
	GetSubstraction()
	GetMultiplication()
	GetDivision()
}
