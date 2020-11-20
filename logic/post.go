package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1.generate ID
	p.PostID = snowflake.GenID()

	return mysql.CreatePost(p)
}
