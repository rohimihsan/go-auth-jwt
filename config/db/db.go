package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbname = "bakulan"

func Con() (*mongo.Database, error) {
	clientOption := options.Client().ApplyURI("mongodb+srv://appUser:Zc7NAnSYoe68VugQ@bakulan.ibqq4.mongodb.net/<dbname>?retryWrites=true&w=majority")

	//initate connection
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	database := client.Database(dbname)

	return database, nil
}
