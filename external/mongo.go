package external

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"lessbutter.co/mealkit/config"
)

func MongoConn() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	conf := config.GetConfiguration()
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	credential := options.Credential{
		Username: conf.MONGO_USERNAME,
		Password: conf.MONGO_PASSWORD,
	}
	clientOptions := options.Client().ApplyURI(conf.MONGO_URL).SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOptions)

	CheckErr(err)
	CheckErr(client.Ping(ctx, readpref.Primary()))

	fmt.Println("MongoDB Connection Made")
	return client, ctx, cancel
}
