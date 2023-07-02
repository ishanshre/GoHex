package ports

type DbPort interface {
	CloseDbConnection()
	AddToHistroy(answer int32, operation string) error
}
