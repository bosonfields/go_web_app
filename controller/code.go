package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "requesting params error",
	CodeUserExist:       "user exist",
	CodeUserNotExist:    "user not exist",
	CodeInvalidPassword: "username or password error",
	CodeServerBusy:      "server busy",

	CodeNeedLogin:    "need login",
	CodeInvalidToken: "invalid token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
