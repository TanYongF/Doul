package rabbitmq

// LikesRelationUpdateStockMessage  点赞消息入库操作
type LikesRelationUpdateStockMessage struct {
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
	Type    bool  `json:"type"`
}

// FollowRelationUpdateMessage  用户关注消息入库操作
type FollowRelationUpdateMessage struct {
	FollowerId  int64 `json:"follower_id"`
	FollowingId int64 `json:"following_id"`
	Type        bool  `json:"type"`
}
