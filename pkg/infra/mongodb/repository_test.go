package mongodb

import (
	"context"
	"errors"
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

func TestUserRepository_GetUser(t *testing.T) {
	repo, err := NewMongoUserRepository(context.TODO(), mongoServer.URI())
	if err != nil {
		t.Error(err)
	}

	// Create a new User
	expectedUser := user.User{
		Email:    "test@test.com",
		Password: "",
		Id:       uuid.New(),
	}
	newUser := Create(expectedUser)
	_, insertErr := repo.user.InsertOne(context.TODO(), newUser)
	if insertErr != nil {
		t.Error(insertErr)
	}

	// Build our needed testcase data struct
	type testCase struct {
		name        string
		arg         uuid.UUID
		expectedErr error
	}
	// Create new test cases
	testCases := []testCase{
		{
			name:        "Get User by id",
			arg:         expectedUser.Id,
			expectedErr: nil,
		},
		{
			name:        "Get User by non-existent id",
			arg:         uuid.New(),
			expectedErr: errors.New("user not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Get User
			_, err := repo.GetUser(tc.arg)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestUserRepository_GetAllUsers(t *testing.T) {
	repo, err := NewMongoUserRepository(context.TODO(), mongoServer.URI())
	if err != nil {
		log.Fatal(err)
	}

	// Create new slice of Users
	expectedUsers := [2]user.User{
		{
			Email: "test1@test.com",
			Id:    uuid.New(),
		},
		{
			Email: "test2@test.com",
			Id:    uuid.New(),
		},
	}
	mongoUsers := make([]interface{}, 2)
	for _, eu := range expectedUsers {
		newUser := Create(eu)
		mongoUsers = append(mongoUsers, newUser)
	}
	_, insertErr := repo.user.InsertMany(context.TODO(), mongoUsers)
	if insertErr != nil {
		t.Error(insertErr)
	}

	users, err := repo.GetAllUsers()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, len(users), "Expected 1 user")
}
