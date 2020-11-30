package mongodb

import (
	"context"
	"fmt"
	"web_app/settings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var client *mongo.Client
var mongodb *mongo.Database

func Init(cfg *settings.MongodbConfig) (err error) {
	// Set client options
	dsn := fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port)
	credential := options.Credential{
		Username: cfg.User,
		Password: cfg.Password,
	}
	clientOptions := options.Client().ApplyURI(dsn).SetAuth(credential)
	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		zap.L().Error("init mongodb error", zap.Error(err))
		return
	}
	mongodb = client.Database("bluebell")

	if err != nil {
		zap.L().Error("connect database error", zap.Error(err))
		return
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		zap.L().Error("ping mongodb error", zap.Error(err))
		return
	}
	zap.L().Info("Connect to MongoDB")
	return
}
func Close() (err error) {
	err = client.Disconnect(context.TODO())
	if err != nil {
		zap.L().Error("close mongodb error", zap.Error(err))
		return
	}
	zap.L().Info("Connection to MongoDB closed.")
	return
}
