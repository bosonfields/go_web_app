package logic

import (
	"web_app/dao/mongodb"
	"web_app/models"
)

func GetEventById(eid int64) (data *models.Event, err error) {
	return mongodb.GetEventById(eid)
}
