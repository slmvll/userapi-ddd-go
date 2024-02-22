package mongodb

import (
	"context"
	"time"

	"userapi-ddd-go/pkg/domain/user"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserRepository struct {
	db   *mongo.Database
	user *mongo.Collection
}

func NewMongoUserRepository(ctx context.Context, connectionString string) (*MongoUserRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("user-db")
	users := db.Collection("user")

	return &MongoUserRepository{
		db:   db,
		user: users,
	}, nil
}

func (ur *MongoUserRepository) AddUser(user user.User) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newUser := Create(user)
	_, err := ur.user.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}
	return nil
}

func (ur *MongoUserRepository) GetUser(id uuid.UUID) (user.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultUser := ur.user.FindOne(ctx, bson.M{"id": id})

	var retrievedUser MongoUser
	err := resultUser.Decode(&retrievedUser)
	if err != nil {
		return user.User{}, err
	}
	return retrievedUser.ToDomain(), nil
}

func (ur *MongoUserRepository) GetAllUsers() ([]user.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cursor, err := ur.user.Find(ctx, bson.D{})

	var resultsUsers []MongoUser
	if err = cursor.All(ctx, &resultsUsers); err != nil {
		return []user.User{}, err
	}

	var retrievedUsers []user.User
	for _, result := range resultsUsers {
		t := result.ToDomain()
		retrievedUsers = append(retrievedUsers, t)
	}

	return retrievedUsers, nil
}

func (m MongoUser) ToDomain() user.User {
	// Create a ProxyUser
	u := user.User{
		Id:       m.Id,
		Email:    m.Email,
		Password: m.Password,
		Status:   m.Status,
	}
	return u
}

func Create(u user.User) *MongoUser {
	m := MongoUser{u.Id, u.Email, u.Password, u.Status, time.Now()}
	return &m
}
