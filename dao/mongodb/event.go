package mongodb

import (
	"context"
	"web_app/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetEventById(eid int64) (event *models.Event, err error) {

	event = new(models.Event)
	filter := bson.D{{"event_id", eid}}

	err = mongodb.Collection("events").FindOne(context.TODO(), filter).Decode(&event)

	if err != nil {
		//zap.L().Error("No event id matched in mongodb", zap.Error(err))
		return
	}
	return
}
