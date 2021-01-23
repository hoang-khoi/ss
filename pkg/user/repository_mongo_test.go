package user

import (
	"context"
	"log"
	"os"
	"ss/container"
	"ss/infra"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoContainer = container.MongoContainer{ContainerName: "mongo_test_container"}
var client *mongo.Client

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestCreateAndFindUser(t *testing.T) {
	purgeDb()
	repo := getRepo()
	originalUser := Model{
		ID:       "koi",
		Password: "secret",
	}

	_ = repo.Create(&originalUser)

	retrievedUser, _ := repo.Find("koi")
	assert.Equal(t, originalUser, *retrievedUser)
}

func TestFindNotFound(t *testing.T) {
	purgeDb()
	repo := getRepo()
	u, _ := repo.Find("black_dad")
	assert.Nil(t, u)
}

// Needs to call this one manually before each test
func purgeDb() {
	if err := client.Database("ss").Drop(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func setup() {
	_ = mongoContainer.Stop()

	if err := mongoContainer.Start(); err != nil {
		log.Fatal(err)
	}

	c, err := infra.MongoFactory{
		URI:           "mongodb://localhost:27017",
		TimeoutSecond: 1,
	}.GetClient()

	if err != nil {
		log.Fatal(err)
	}

	client = c
}

func teardown() {
	_ = client.Disconnect(context.Background())
	if err := mongoContainer.Stop(); err != nil {
		log.Fatal(err)
	}
}

func getRepo() RepositoryMongo {
	return RepositoryMongo{Client: client}
}
