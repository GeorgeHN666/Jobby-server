package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB_NAME = "jobby"

func (app *application) connect() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(app.config.db.URI))
	if err != nil {
		app.errorLog.Fatal(err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		app.errorLog.Fatal(err.Error())
	}

	app.infoLog.Printf("Mongo Db connected and redy to use in PORT %d", app.config.port)
	return client
}

// USER SECTION
func (app *application) InsertUser(u User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := app.connect().Database(DB_NAME)
	col := db.Collection("users")

	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.IsAuth = "false"
	u.Role = 1
	u.New = "true"

	_, err := col.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	return nil
}
