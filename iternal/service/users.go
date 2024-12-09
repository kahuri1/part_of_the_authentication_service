package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/domain"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/model"
)

const (
	verificationCodeLength = 6
)

func (s *Service) SignUpService(ctx context.Context, input model.UserSignUpInput) error {
	passwordHash, err := s.passwordHasher.Hash(input.Password)
	if err != nil {
		return err
	}

	verificationCode := s.otpGenerator.RandomSecret(verificationCodeLength)

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: passwordHash,
		Verification: model.Verification{
			Code: verificationCode,
		},
	}
	if err = s.repo.CreateRepo(ctx, user); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return err
		}

		return err
	}
	go func() {
		err = s.emailSender.SendVerificationCode(user.Email, verificationCode)
		if err != nil {
			fmt.Println("email sending verification code error:", err)
		}
	}()

	return nil
}
