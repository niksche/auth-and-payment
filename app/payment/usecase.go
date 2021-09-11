package payment

type UseCase interface {
	MakePayment(username string) error
}
