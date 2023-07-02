package api

import (
	"github.com/ishanshre/GoHex/internal/ports"
)

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{
		db:    db,
		arith: arith,
	}
}

func (apia Adapter) GetAddition(a, b int) (int, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	if err := apia.db.AddToHistroy(int32(answer), "addition"); err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetSubstraction(a, b int) (int, error) {
	answer, err := apia.arith.Substraction(a, b)
	if err != nil {
		return 0, err
	}
	if err := apia.db.AddToHistroy(int32(answer), "substraction"); err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetMultiplication(a, b int) (int, error) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	if err := apia.db.AddToHistroy(int32(answer), "multiplication"); err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetDivision(a, b int) (int, error) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	if err := apia.db.AddToHistroy(int32(answer), "division"); err != nil {
		return 0, err
	}
	return answer, nil
}
