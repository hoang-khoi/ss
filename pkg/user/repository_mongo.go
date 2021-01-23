package user

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

// DB Database's name for this repository.
const DB = "ss"

// CollectionName Mongo collection for users.
const CollectionName = "users"

const timeoutSecond = 2

// RepositoryMongo implements UserRepository with MongoDB.
type RepositoryMongo struct {
	Client *mongo.Client
}

//Create creates a new user.
func (m *RepositoryMongo) Create(u *Model) (e error) {
	ctx, cancel := getContext()
	defer cancel()

	collection := getUsersCollection(m.Client)
	_, e = collection.InsertOne(ctx, u)

	return
}

// Find retrieves the user by id, returns nil if not found.
func (m *RepositoryMongo) Find(id string) (*Model, error) {
	collection := getUsersCollection(m.Client)
	var user Model
	ctx, cancel := getContext()
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	return &user, err
}

func getUsersCollection(c *mongo.Client) *mongo.Collection {
	return c.Database(DB).Collection(CollectionName)
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*timeoutSecond)
}
