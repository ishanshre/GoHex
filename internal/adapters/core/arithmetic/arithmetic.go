package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith Adapter) Addition(a int, b int) (int, error) {
	return a + b, nil
}

func (arith Adapter) Substraction(a int, b int) (int, error) {
	return a - b, nil
}

func (arith Adapter) Multiplication(a, b int) (int, error) {
	return a * b, nil
}

func (arith Adapter) Division(a, b int) (int, error) {
	if b == 0 {
		panic("b cannot be zero")
	}
	return a / b, nil
}
