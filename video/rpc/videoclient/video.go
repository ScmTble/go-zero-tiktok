// Code generated by goctl. DO NOT EDIT!
// Source: video.proto

package videoclient

import (
	"context"
	video2 "tiktok/video/rpc/video"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PublishVideo         = video2.PublishVideo
	PublishVideoListReq  = video2.PublishVideoListReq
	PublishVideoListResp = video2.PublishVideoListResp
	PublishVideoReq      = video2.PublishVideoReq
	PublishVideoResp     = video2.PublishVideoResp

	Video interface {
		PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*PublishVideoResp, error)
		PublishVideoList(ctx context.Context, in *PublishVideoListReq, opts ...grpc.CallOption) (*PublishVideoListResp, error)
	}

	defaultVideo struct {
		cli zrpc.Client
	}
)

func NewVideo(cli zrpc.Client) Video {
	return &defaultVideo{
		cli: cli,
	}
}

func (m *defaultVideo) PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*PublishVideoResp, error) {
	client := video2.NewVideoClient(m.cli.Conn())
	return client.PublishVideo(ctx, in, opts...)
}

func (m *defaultVideo) PublishVideoList(ctx context.Context, in *PublishVideoListReq, opts ...grpc.CallOption) (*PublishVideoListResp, error) {
	client := video2.NewVideoClient(m.cli.Conn())
	return client.PublishVideoList(ctx, in, opts...)
}
