package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go_code/Doul/app/video/cmd/rpc/internal/svc"
	"go_code/Doul/app/video/cmd/rpc/video"
	"go_code/Doul/common/globalkey"
	"go_code/Doul/common/rabbitmq"
	"go_code/Doul/common/xerr"
	"strconv"
)

type LikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const luaScript = `
	local change = tonumber(ARGV[3])
	local exist = redis.call('SISMEMBER', KEYS[1], ARGV[1]) 
	if change == 1 then
		if exist == 0 then
			redis.call("HINCRBY", KEYS[2], ARGV[2], change)
			redis.call("SADD", KEYS[1], ARGV[1])
		end	
	else
		redis.call("HINCRBY", KEYS[2], ARGV[2], change)
		redis.call("Srem", KEYS[1], ARGV[1])
	end
	return 1
`

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeLogic) Like(in *video.LikeReq) (*video.LikeResp, error) {
	var err error
	//todo 如何解决重复点赞问题？
	// step1 : update the counter
	exist, err := l.svcCtx.RedisClient.Hexists(globalkey.GetVideoLikesCounterRedisKey(in.VideoId),
		globalkey.GetVideoLikesCounterFieldKey(in.VideoId))
	if err != nil {
		return nil, err
	}
	if !exist {
		l.svcCtx.DyFavoriteModel.CountLikesByVideoId(l.ctx, in.VideoId)
	}

	var change int
	if in.Type {
		change = 1
	} else {
		change = -1
	}
	_, err = l.svcCtx.RedisClient.EvalCtx(l.ctx, luaScript, []string{
		globalkey.GetVideoLikesUsersRedisKey(in.VideoId),
		globalkey.GetVideoLikesCounterRedisKey(in.VideoId),
	}, []string{
		strconv.FormatInt(in.UserId, 10),
		globalkey.GetVideoLikesCounterFieldKey(in.VideoId),
		strconv.Itoa(change),
	})
	if err != nil {
		return nil, err
	}
	//// 判断点赞类型
	//if in.Type {
	//	err = l.svcCtx.RedisClient.PipelinedCtx(l.ctx, func(pipe redis.Pipeliner) error {
	//		_, err = l.svcCtx.RedisClient.Hincrby(
	//			globalkey.GetVideoLikesCounterRedisKey(in.VideoId),
	//			globalkey.GetVideoLikesCounterFieldKey(in.VideoId),
	//			1)
	//		_, err = l.svcCtx.RedisClient.Sadd(globalkey.GetVideoLikesUsersRedisKey(in.VideoId), in.UserId)
	//		return err
	//	})
	//
	//} else {
	//	err = l.svcCtx.RedisClient.PipelinedCtx(l.ctx, func(pipe redis.Pipeliner) error {
	//		_, err = l.svcCtx.RedisClient.Hincrby(
	//			globalkey.GetVideoLikesCounterRedisKey(in.VideoId),
	//			globalkey.GetVideoLikesCounterFieldKey(in.VideoId),
	//			-1)
	//		_, err = l.svcCtx.RedisClient.Srem(globalkey.GetVideoLikesUsersRedisKey(in.VideoId), in.UserId)
	//		return err
	//	})
	//}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CACHE_ERROR), "update the %d video counter error", in.VideoId)
	}

	// step2 : send message to mq
	// action : 1: like 2: dislike
	message, err := json.Marshal(rabbitmq.LikesRelationUpdateStockMessage{
		UserId:  in.UserId,
		VideoId: in.VideoId,
		Type:    in.Type,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Error"), "marshal error")
	}

	//异步入库操作
	err = l.svcCtx.MqSender.Send("likes", "", message)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.MQ_ERROR), "public message error")
	}

	return &video.LikeResp{}, nil
}