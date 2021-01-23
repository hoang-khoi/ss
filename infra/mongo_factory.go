package infra

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoFactory simplifies the ugly mongo client creation code.
type MongoFactory struct {
	URI           string
	TimeoutSecond time.Duration
}

// GetClient build a client to mongo. Method Disconnect must be call to release the resources.
func (m MongoFactory) GetClient() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.URI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), m.TimeoutSecond*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
