package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在

	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// generate UID
	userID := snowflake.GenID()

	// struct a user instance

	//fmt.Println(userID)

	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// password encipher
	// save to database

	// redis.xxx

	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}

	// generate JWT
	zap.L().Info("check login userID in login", zap.Int64("p.userID", user.UserID))
	return jwt.GenToken(user.UserID, user.Username)
}
