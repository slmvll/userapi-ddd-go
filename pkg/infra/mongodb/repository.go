package mongodb

import (
	"context"
	"time"

	"userapi-ddd-go/pkg/domain/user"

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

func Create(u user.User) *MongoUser {
	m := MongoUser{u.Id, u.Email, u.Password, u.Status, time.Now()}
	return &m
}
