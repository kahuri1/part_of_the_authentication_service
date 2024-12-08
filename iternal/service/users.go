package service

import (
	"context"
	"errors"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/domain"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/model"
)

const (
	verificationCodeLength = 6
)

func (s *Service) SignUp(ctx context.Context, input model.UserSignUpInput) error {
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
	if err = s.repo.Create(ctx, user); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return err
		}

		return err
	}

	return s.emailSender.SendVerificationCode(user.Email, verificationCode)
}
