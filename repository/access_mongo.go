package repository

import (
	"context"
	"time"

	"github.com/blackestwhite/zwrapper/entity"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAccessTokenRepository struct {
	collection *mongo.Collection
}

func NewMongoAccessTokenRepository(client *mongo.Client, databaseName, collectionName string) *MongoAccessTokenRepository {
	collection := client.Database(databaseName).Collection(collectionName)
	return &MongoAccessTokenRepository{
		collection: collection,
	}
}

func (repo *MongoAccessTokenRepository) CreateToken(accessToken entity.AccessToken) (inserted entity.AccessToken, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	u, err := uuid.NewV4()
	if err != nil {
		return
	}
	accessToken.Token = u.String()

	_, err = repo.collection.InsertOne(ctx, accessToken)
	return accessToken, err
}

func (repo *MongoAccessTokenRepository) GetByToken(token string) (accessToken entity.AccessToken, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = repo.collection.FindOne(ctx, bson.M{
		"token": token,
	}).Decode(&accessToken)
	return
}
