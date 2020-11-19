package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)

// 把每一步数据库操作封装

const secret = "bosonfields.com"

var (
	ErrorUserExist       = errors.New("user already exist")
	ErrorUserNotExist    = errors.New("user not exist")
	ErrorInvalidPassword = errors.New("incorrect password")
)

// CheckUserExist
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`

	var count int
	//if err := db.Get(&count, sqlStr, username); err != nil {
	//	return err
	//}
	err = db.Get(&count, sqlStr, username)
	if count > 0 {
		return ErrorUserExist
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

func Login(user *models.User) (err error) {
	oPassword := user.Password

	sqlStr := `select user_id, username, password from user where username = ?`

	err = db.Get(user, sqlStr, user.Username)

	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	// judge password correct

	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
