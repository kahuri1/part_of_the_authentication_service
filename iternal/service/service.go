package service

import (
	"context"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/model"
	"github.com/kahuri1/part_of_the_authentication_service/pkg/email"
	"github.com/kahuri1/part_of_the_authentication_service/pkg/otp"
	log "github.com/sirupsen/logrus"
)

type Repo interface {
	CreateRepo(ctx context.Context, user model.User) error
	CheckUserRepo(uuid string) error
	CheckSessionByUuidUserRepo(uuid string) (model.RefreshSession, error)
	CreateSessionRepo(auth model.AuthenticationRequest, refreshToken []byte) error
	GetRefreshSessionByRefreshTokenRepo(token model.Tokens) (model.RefreshSession, error)
	UpdateRefreshSessionRepo(uuid string, ip string, refreshToken []byte) error
}
type passwordHasher interface {
	Hash(password string) (string, error)
}

type Service struct {
	repo           Repo
	passwordHasher passwordHasher
	otpGenerator   otp.Generator
	emailSender    email.EmailSender
}

func NewService(repo Repo, passwordHasher passwordHasher, otpGenerator otp.Generator, emailSender email.EmailSender) *Service {
	log.Info("service init")

	return &Service{
		repo:           repo,
		passwordHasher: passwordHasher,
		otpGenerator:   otpGenerator,
		emailSender:    emailSender,
	}
}
