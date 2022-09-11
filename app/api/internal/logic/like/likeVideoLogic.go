package like

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"tiktok/app/api/internal/svc"
	"tiktok/app/api/internal/types"
	"tiktok/config"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LikeVideoLogic) LikeVideo(req *types.LikeReq) (resp *types.LikeResp, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	key := strings.Builder{}
	key.WriteString(config.LsikeVideoRedisKey)
	key.WriteString(strconv.FormatInt(req.VideoId, 10))
	r, err := l.svcCtx.Redis.Sadd(key.String(), uid)
	err = l.svcCtx.Redis.Expire(key.String(), 640000)
	if err != nil {
		// 点赞设置过期时间出错，删除key
		l.svcCtx.Redis.Del(key.String())
		return nil, err
	}
	fmt.Println(r)
	if err != nil {
		return nil, err
	}
	return &types.LikeResp{
		StatusCode: true,
	}, nil
}
