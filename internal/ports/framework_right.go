package ports

type DbPort interface {
	CloseDbConnection()
	AddToHistory(answer int, operation string) error
}
