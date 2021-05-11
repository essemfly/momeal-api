package external

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"lessbutter.co/mealkit/config"
	"lessbutter.co/mealkit/utils"
)

func MongoConn() (client *mongo.Client) {
	conf := config.GetConfiguration()
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
	return client
}
