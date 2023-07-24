package repository

import (
	"context"
	"time"

	"github.com/blackestwhite/zwrapper/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPaymentRepository struct {
	collection *mongo.Collection
}

func NewMongoPaymentRepository(client *mongo.Client, databaseName, collectionName string) *MongoPaymentRepository {
	collection := client.Database(databaseName).Collection(collectionName)
	return &MongoPaymentRepository{
		collection: collection,
	}
}

func (repo *MongoPaymentRepository) CreatePayment(payment entity.Payment) (inserted entity.Payment, err error) {
	payment.ID = primitive.NewObjectID()
	payment.Timestamp = time.Now().Unix()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = repo.collection.InsertOne(ctx, payment)
	return payment, err
}

func (repo *MongoPaymentRepository) GetPayment(id string) (payment entity.Payment, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = repo.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&payment)
	return
}
