package mongo

import (
	"context"
	"mini-wallet/src/helpers"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context) (*mongo.Database, error) {
	db := helpers.GetEnv("DB_URI", "")
	dbTimeout := helpers.GetEnv("DB_TIMEOUT", "10000")

	clientOptions := options.Client().ApplyURI(db)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	timeout, _ := strconv.Atoi(dbTimeout)
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(helpers.GetEnv("DB_NAME", "")), err
}
