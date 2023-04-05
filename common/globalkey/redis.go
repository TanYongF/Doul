package globalkey

import (
	"strconv"
	"time"
)

var (
	TokenPrefix     = "cache:douyin:dyUser:token:"
	UserPrefix      = "user:"
	TokenExpireTime = time.Hour * 24
)

func GetVideoLikesCounterRedisKey(videoId int64) string {
	return "doul:counter:video:likes:" + strconv.FormatInt(videoId>>9, 10)
}
func GetVideoLikesCounterFieldKey(videoId int64) string {
	return strconv.FormatInt(videoId&0x1FF, 10)
}

func GetVideoCommentsCounterKey(videoId int64) string {
	return "doul:counter:video:comments:" + strconv.FormatInt(videoId>>9, 10)
}
func GetVideoCommentsCounterFieldKey(videoId int64) string {
	return strconv.FormatInt(videoId&0x1FF, 10)
}

func GetVideoLikesUsersRedisKey(videoId int64) string {
	return "doul:video:" + strconv.FormatInt(videoId, 10) + ":recently"
}

func GetUserByToken(token string) string {
	return "doul:token:" + token
}

func GetUserById(userId int64) string {
	return "doul:user:userId:" + strconv.FormatInt(userId, 10)
}
