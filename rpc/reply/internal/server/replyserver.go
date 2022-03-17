// Code generated by taoctl. DO NOT EDIT!
// Source: reply.proto

package server

import (
	"context"

	"sllt/backend-learning/blog/rpc/reply/internal/logic"
	"sllt/backend-learning/blog/rpc/reply/internal/svc"
	"sllt/backend-learning/blog/rpc/reply/reply"
)

type ReplyServer struct {
	svcCtx *svc.ServiceContext
}

func NewReplyServer(svcCtx *svc.ServiceContext) *ReplyServer {
	return &ReplyServer{
		svcCtx: svcCtx,
	}
}

func (s *ReplyServer) CreateReply(ctx context.Context, in *reply.CreateReplyReq) (*reply.CreateReplyResp, error) {
	l := logic.NewCreateReplyLogic(ctx, s.svcCtx)
	return l.CreateReply(in)
}

func (s *ReplyServer) GetReplies(ctx context.Context, in *reply.GetRepliesReq) (*reply.GetRepliesResp, error) {
	l := logic.NewGetRepliesLogic(ctx, s.svcCtx)
	return l.GetReplies(in)
}