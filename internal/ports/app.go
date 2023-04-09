package ports

type ApiPort interface {
	GetAddition(a, b int) (int, error)
	GetSubstraction(a, b int) (int, error)
	GetMultiplication(a, b int) (int, error)
	GetDivision(a, b int) (int, error)
}
