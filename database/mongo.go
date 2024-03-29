package database

import (
	"context"
	"time"

	"github.com/lessbutter/momeal-api/config"
	"github.com/lessbutter/momeal-api/src/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Db *mongo.Database

func InitDB(conf config.Configuration) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	credential := options.Credential{
		Username: conf.MONGO_USERNAME,
		Password: conf.MONGO_PASSWORD,
	}
	clientOptions := options.Client().ApplyURI(conf.MONGO_URL).SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOptions)

	utils.CheckErr(err)
	utils.CheckErr(client.Ping(ctx, readpref.Primary()))
	Db = client.Database("mealkit_new")
}
