package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = func() context.Context {
	return context.Background()
}()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connect() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}

func aggregate() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	pipeline := make([]bson.M, 0)
	err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
		[
			{ "$group": {
				"_id": null,
				"Total": { "$sum": 1 }
			} },
			{ "$project": {
				"Total": 1,
				"_id": 0
			} }
		]
	`)), true, &pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Total :", result[0]["Total"])
	}
}

func main() {
	aggregate()
}
