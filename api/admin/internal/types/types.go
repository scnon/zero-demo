// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type AddUserReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	UserName string `json:"userName"`
	Status   int    `json:"status,default=1"`
	Sort     int    `json:"sort,default=0"`
	Remark   string `json:"remark,default=''"`
}

type DeleteUserReq struct {
	Ids []int `json:"ids"`
}

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginResp struct {
	UserName     string `json:"userName"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireTime   int64  `json:"expireTime"`
}

type QueryUserListReq struct {
	Account  string `json:"account,optional"`
	UserName string `json:"userName,optional"`
	Status   int    `json:"status,default=1"`
	Page     int    `json:"page,default=1"`
	PageSize int    `json:"pageSize,default=10"`
}

type QueryUserListResp struct {
	Total int        `json:"total"`
	List  []UserData `json:"list"`
}

type UpdateUserReq struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	UserName string `json:"userName"`
	Status   int    `json:"status"`
	Sort     int    `json:"sort"`
	Remark   string `json:"remark"`
}

type UserData struct {
	Id         int    `json:"id"`
	Account    string `json:"account"`
	UserName   string `json:"userName"`
	Status     int    `json:"status"`
	Sort       int    `json:"sort"`
	Remark     string `json:"remark"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
