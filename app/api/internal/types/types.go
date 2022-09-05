// Code generated by goctl. DO NOT EDIT.
package types

type RegisterRequest struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type RegisterResponse struct {
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type LoginResponse struct {
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	UserId         int64  `json:"userId"`
	UserName       string `json:"userName"`
	FollowCount    int64  `json:"followCount"`
	FollowerCount  int64  `json:"followerCount"`
	IsFollow       bool   `json:"isFollow"`
	TotalFavorited int64  `json:"totalFavorited"`
	FavoriteCount  int64  `json:"favoriteCount"`
}

type PublishVideoReq struct {
	PlayUrl  string `json:"playUrl"`
	CoverUrl string `json:"coverUrl"`
	Title    string `json:"title"`
}

type PublishVideoResp struct {
	PlayUrl  string `json:"playUrl"`
	CoverUrl string `json:"coverUrl"`
	Title    string `json:"title"`
}