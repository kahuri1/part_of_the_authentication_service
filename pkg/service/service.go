package service

import log "github.com/sirupsen/logrus"

type Repo interface {
}

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	log.Info("service init")

	return &Service{
		repo: repo,
	}
}
