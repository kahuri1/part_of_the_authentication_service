package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/domain"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/model"
	"github.com/spf13/viper"
	"time"
)

func (s *Service) AuthenticationService(auth model.AuthenticationRequest) (model.Tokens, error) {
	if len(auth.Uuid) == 0 {
		return model.Tokens{}, errors.New("uuid is empty")
	}

	err := s.repo.CheckUserRepo(auth.Uuid)
	if err != nil {
		return model.Tokens{}, err
	}

	accessToken, err := createJwt(auth)
	if err != nil {
		return model.Tokens{}, err
	}
	refreshToken, err := createRefreshToken()
	if err != nil {
		return model.Tokens{}, err
	}

	dataSession, err := s.repo.CheckSessionByUuidUserRepo(auth.Uuid)
	if errors.Is(err, domain.ErrSessionOpen) {
		err = s.repo.UpdateRefreshSessionRepo(dataSession.UserUuid, auth.Ip, refreshToken)
		return model.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken}, nil
	}

	err = s.repo.CreateSessionRepo(auth, refreshToken)
	if err != nil {
		return model.Tokens{}, err
	}

	return model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken}, nil
}

func createJwt(auth model.AuthenticationRequest) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": auth.Uuid, // Subject (user identifier)
		"ip":  auth.Ip,
		"iss": "todo-app",                                               // Issuer
		"exp": time.Now().Add(viper.GetDuration("auth.accessTokenTTL")), // Expiration time
		"iat": time.Now().Unix(),                                        // Issued at
	})

	return claims.SignedString([]byte(viper.GetString("key.secretKey")))
}

func createRefreshToken() (string, error) {
	newUUID := uuid.New()
	return newUUID.String(), nil
}

func (s *Service) RefreshTokenService(Token model.Tokens, ip string) (model.Tokens, error) {

	if len(Token.RefreshToken) == 0 {
		return model.Tokens{}, errors.New("refresh Token is empty")
	}

	dataSession, err := s.repo.GetRefreshSessionByRefreshTokenRepo(Token)
	if err != nil {
		return model.Tokens{}, err
	}
	if isExpired(dataSession.ExpiresAt) {
		return model.Tokens{}, errors.New("refresh token expired, log in again")
	}
	accessToken, err := createJwt(model.AuthenticationRequest{
		Uuid: dataSession.UserUuid,
		Ip:   ip,
	})
	if err != nil {
		return model.Tokens{}, err
	}
	refreshToken, err := createRefreshToken()
	if err != nil {
		return model.Tokens{}, err
	}

	if ip != dataSession.Ip {
		go func() {
			err = s.emailSender.WarningMessageIP(ip)
			if err != nil {
				fmt.Printf("Failed to send warning email: %v\n", err)
			}
		}()
	}
	err = s.repo.UpdateRefreshSessionRepo(dataSession.UserUuid, ip, refreshToken)
	if err != nil {
		return model.Tokens{}, err
	}

	return model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken}, nil
}

func isExpired(expiresAt time.Time) bool {
	return time.Now().After(expiresAt)
}