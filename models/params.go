package models

// define params struct

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    int64 `json:"post_id,string" binding:"required"`
	Direction int8  `json:"direction,string" binding:"required,oneof=1 0 -1"` // agree 1, disagree -1, cancel 0
}
