package usecase

import (
	"fmt"

	"github.com/niksche/flex/app/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo auth.UserRepository
}

func NewAuthUseCase(userRepo auth.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
	}
}

func (u *AuthUseCase) SignUp(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Errorf("bcrypt function fails")
		return err
	}

	if err := u.userRepo.CreateUser(username, string(hashedPassword)); err != nil {
		fmt.Errorf("cannot sign up user")
		return err
	}
	return nil
}

func (u *AuthUseCase) LoginUser(username, password string) error {
	user, err := u.userRepo.GetUser(username)
	if err != nil {
		return fmt.Errorf("cannot find user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
		fmt.Printf("successfuly login")
	}
	return nil

}

func (u *AuthUseCase) MakePayment(username string) error {
	err := u.userRepo.MakePayment(username)
	if err != nil {
		return fmt.Errorf("cannot find user")
	}
	return nil
}
