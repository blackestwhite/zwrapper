package service

import (
	"context"
	"time"

	"github.com/blackestwhite/zwrapper/db"
	"github.com/blackestwhite/zwrapper/entity"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type AccessTokenService struct{}

func (a *AccessTokenService) Create(instance entity.AccessToken) (at entity.AccessToken, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	u, err := uuid.NewV4()
	if err != nil {
		return
	}
	instance.Token = u.String()

	_, err = db.Client.Database("zwrapper").Collection("tokens").InsertOne(ctx, instance)
	return instance, err
}

func (a *AccessTokenService) GetByToken(token string) (at entity.AccessToken, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = db.Client.Database("zwrapper").Collection("tokens").FindOne(ctx, bson.M{
		"token": token,
	}).Decode(&at)
	return
}
