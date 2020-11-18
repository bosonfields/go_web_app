package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// 判断用户存不存在
	mysql.QueryUserByUsername()
	// generate UID
	snowflake.GenID()
	// password encipher
	// save to database
	mysql.InsertUser()
	// redis.xxx

}
