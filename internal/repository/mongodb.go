package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Todo: URL to connect
// mongodb+srv://emprestaai:<password>@cluster0.fzjem4n.mongodb.net/?retryWrites=true&w=majority
func SetupDb() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}
	client := NewDBClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Errorf("error while trying to connect to db %v", err)
	}
	defer client.Disconnect(ctx)

	// debug only
	listaDBs(client, ctx)
}

func NewDBClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://emprestaai:" + os.Getenv("password") + "@cluster0.fzjem4n.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Errorf("error loading database %v", err)
	}
	return client
}

// listaDBs : list databases - debug/dev only feature, its not suppose to work public
func listaDBs(c *mongo.Client, ctx context.Context) {
	databases, err := c.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("list dbs: ", databases)
}
