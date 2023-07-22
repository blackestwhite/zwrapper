package entity

type AccessToken struct {
	Consumer string `bson:"consumer" json:"consumer"`
	Token    string `bson:"token" json:"token"`
}
