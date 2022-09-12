package like

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"
	"tiktok/config"
)

type LikeVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeVideoLogic {
	return &LikeVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// LikeVideo 点赞视频
func (l *LikeVideoLogic) LikeVideo(req *types.LikeReq) (resp *types.LikeResp, err error) {
	if req.StatusCode != config.UserDisLikeVideo && req.StatusCode != config.UserLikeVideo {
		// 参数校验，点赞状态不是0或1
		return nil, errors.New("点赞状态异常")
	}

	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	key := strings.Builder{}
	key.WriteString(config.LsikeVideoRedisKey)
	key.WriteString(strconv.FormatInt(req.VideoId, 10))

	// 从redis中查询是否已经存在数据
	exist, err := l.svcCtx.Redis.SismemberCtx(l.ctx, key.String(), uid)
	if err != nil {
		return nil, err
	}

	if !exist {
		data := make(map[string]int64, 3)
		data["UserId"] = uid
		data["VideoId"] = req.VideoId
		data["StatusCode"] = req.StatusCode
		b, _ := json.Marshal(data)
		if req.StatusCode == config.UserLikeVideo {
			// 不存在且点赞
			// 通过RabbitMQ向rpc发送点赞消息
			if err := l.svcCtx.RabbitMQ.Channel.Publish("LikeExchange", "", false, false, amqp.Publishing{
				ContentType: l.svcCtx.RabbitMQ.ContentType,
				Body:        b,
			}); err != nil {
				return nil, err
			}
			// 向redis中添加数据
			if _, err := l.svcCtx.Redis.Sadd(key.String(), uid); err != nil {
				return nil, err
			}
			if err := l.svcCtx.Redis.Expire(key.String(), config.LikeVideoExpireTime); err != nil {
				// 点赞设置过期时间出错，删除key
				l.svcCtx.Redis.Del(key.String())
				return nil, err
			}
			return &types.LikeResp{
				StatusCode: true,
			}, nil
		}
		// redis中不存在，但取消点赞可能redis过期了，但可能数据库中有
		// 通过RabbitMQ向rpc发送点赞消息
		if err := l.svcCtx.RabbitMQ.Channel.Publish("LikeExchange", "", false, false, amqp.Publishing{
			ContentType: l.svcCtx.RabbitMQ.ContentType,
			Body:        b,
		}); err != nil {
			return nil, err
		}
	}

	if req.StatusCode == config.UserLikeVideo {
		// 存在并且点赞
		// 更新key的过期时间直接返回结果
		//err = l.svcCtx.Redis.Expire(key.String(), config.LikeVideoExpireTime)
		return &types.LikeResp{
			StatusCode: true,
		}, nil
	} else {
		// 存在但取消点赞
		if _, err := l.svcCtx.Redis.Srem(key.String(), uid); err != nil {
			return nil, err
		}
		return &types.LikeResp{
			StatusCode: true,
		}, nil
	}

}
