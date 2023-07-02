package ports

type ArithmeticPort interface {
	Addition(a int, b int) (int, error)
	Substraction(a int, b int) (int, error)
	Multiplication(a int, b int) (int, error)
	Division(a int, b int) (int, error)
}
