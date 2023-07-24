package service

import (
	"github.com/blackestwhite/zwrapper/entity"
	"github.com/blackestwhite/zwrapper/repository"
)

type AccessTokenService struct {
	accessTokenRepository repository.AccessTokenRepository
}

func NewAccessTokenService(accessTokenRepo repository.AccessTokenRepository) *AccessTokenService {
	return &AccessTokenService{
		accessTokenRepository: accessTokenRepo,
	}
}

func (a *AccessTokenService) Create(instance entity.AccessToken) (at entity.AccessToken, err error) {
	return a.accessTokenRepository.CreateToken(instance)
}

func (a *AccessTokenService) GetByToken(token string) (at entity.AccessToken, err error) {
	return a.accessTokenRepository.GetByToken(token)
}
