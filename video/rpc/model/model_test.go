package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"testing"
	"tiktok/video/rpc/internal/config"
)

func TestDefaultVideosModel_FindListByUid(t *testing.T) {
	var c config.Config
	conf.MustLoad("../etc/video.yaml", &c)
	videoModel := NewVideosModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis)
	videos, err := videoModel.FindListByUid(context.Background(), 20044)
	fmt.Println(videos)
	fmt.Println(err)
}
