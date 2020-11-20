package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//get all the community and return

	return mysql.GetCommunityList()
}

func GetCommunityDetailByID(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
