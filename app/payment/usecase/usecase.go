package usecase

import (
	"fmt"

	"github.com/niksche/flex/app/payment"
)

type PaymentUseCase struct {
	paymentRepo payment.UserRepository
}

func NewPaymentUseCase(paymentRepo payment.UserRepository) *PaymentUseCase {
	return &PaymentUseCase{
		paymentRepo: paymentRepo,
	}
}

func (u *PaymentUseCase) MakePayment(username string) error {
	err := u.paymentRepo.MakePayment(username)
	if err != nil {
		return fmt.Errorf("cannot find user")
	}
	return nil
}
