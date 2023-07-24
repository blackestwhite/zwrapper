package repository

import "github.com/blackestwhite/zwrapper/entity"

type AccessTokenRepository interface {
	CreateToken(accessToken entity.AccessToken) (entity.AccessToken, error)
	GetByToken(token string) (entity.AccessToken, error)
}
