package mongodb

import (
	"time"

	"github.com/google/uuid"
)

type MongoUser struct {
	Id       uuid.UUID `bson:"id"`
	Email    string    `bson:"email"`
	Password string    `bson:"password"`
	Status   bool      `bson:"status"`
	Created  time.Time `bson:"created"`
}
