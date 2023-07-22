package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Amount      int                `bson:"amount" json:"amount"`
	Authority   string             `bson:"authority" json:"authority"`
	Next        string             `bson:"next" json:"next"`
	Webhook     string             `bson:"webhook" json:"webhook"`
	Key         string             `bson:"key" json:"key"`
	Description string             `bson:"description" json:"description"`
	Timestamp   int64              `bson:"timestamp" json:"timestamp"`
}
