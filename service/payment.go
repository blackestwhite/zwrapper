package service

import (
	"context"
	"time"

	"github.com/blackestwhite/zwrapper/db"
	"github.com/blackestwhite/zwrapper/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentService struct{}

func (p *PaymentService) Create(instance entity.Payment) (inserted entity.Payment, err error) {
	instance.ID = primitive.NewObjectID()
	instance.Timestamp = time.Now().Unix()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = db.Client.Database("zwrapper").Collection("payments").InsertOne(ctx, instance)
	return instance, err
}

func (p *PaymentService) Get(id string) (payment entity.Payment, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = db.Client.Database("zwrapper").Collection("payments").FindOne(ctx, bson.M{"_id": objectID}).Decode(&payment)
	return
}
