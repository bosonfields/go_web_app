package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"web_app/models"
)

// 把每一步数据库操作封装

const secret = "bosonfields.com"

// CheckUserExist
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`

	var count int
	//if err := db.Get(&count, sqlStr, username); err != nil {
	//	return err
	//}
	err = db.Get(&count, sqlStr, username)
	if count > 0 {
		return errors.New("user already exit")
	}

	return
}

// InsertUser
func InsertUser(user *models.User) (err error) {
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	//encipher password
	user.Password = encryptPassword(user.Password)
	//into sql
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
