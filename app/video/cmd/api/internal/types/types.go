// Code generated by goctl. DO NOT EDIT.
package types

type CommonResp struct {
	Code uint32 `json:"status_code"`
	Msg  string `json:"status_msg"`
}

type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	VideoId       int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	UserId        int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

type FeedReq struct {
	LastestTime string `form:"latest_time"`
}

type FeedResp struct {
	CommonResp
	NextTime  *int64  `json:"next_time"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time\
	VideoList []Video `json:"video_list"` // 视频列表
}

type PublishActionReq struct {
	Title string `json:"title"`
}

type PublishActionResp struct {
	CommonResp
}

type PublishListReq struct {
	UserId string `form:"user_id"`
}

type PublishListResp struct {
	CommonResp
	VideoList []Video `json:"video_list"` // 用户发布的视频列表
}

type FavoriteVideofListReq struct {
	UserId string `form:"user_id"`
}

type FavoriteVideoListResp struct {
	CommonResp
	VideoList []Video `json:"video_list"` // 用户喜欢的视频列表
}
