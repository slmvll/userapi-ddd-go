package mongodb

import (
	"context"
	"log"
	"os"
	"testing"
	"userapi-ddd-go/pkg/domain/user"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/bson"
)

var mongoServer *memongo.Server

func TestMain(m *testing.M) {
	var err error
	mongoServer, err = memongo.StartWithOptions(&memongo.Options{MongoVersion: "5.0.20", DownloadURL: "https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-rhel80-5.0.20.tgz"})
	if err != nil {
		log.Fatal(err)
	}
	defer mongoServer.Stop()

	os.Exit(m.Run())
}

func TestUserRepository_AddUser(t *testing.T) {
	repo, err := NewMongoUserRepository(context.TODO(), mongoServer.URI())
	if err != nil {
		log.Fatal(err)
	}

	// Create a new User
	user := user.User{
		Email:    "test@test.com",
		Password: "",
		Id:       uuid.New(),
	}

	addUserErr := repo.AddUser(user)
	if addUserErr != nil {
		t.Error(addUserErr)
	}

	var retrievedUser MongoUser
	resultUsers := repo.user.FindOne(context.TODO(), bson.M{"id": user.Id})
	userErr := resultUsers.Decode(&retrievedUser)
	assert.Nil(t, userErr)
	assert.EqualValues(t, user.Id, retrievedUser.Id)
	assert.EqualValues(t, user.Email, retrievedUser.Email)
}
