package data

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID       string    `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string    `bson:"name" json:"name"`
	Data     string    `bson:"data" json:"data"`
	CreateAt time.Time `bson:"create_at" json:"create_at"`
	UpdateAt time.Time `bson:"update_at" json:"updated_at"`
}
