package logic

import (
	"context"

	"sllt/backend-learning/blog/rpc/reply/internal/svc"
	"sllt/backend-learning/blog/rpc/reply/reply"

	"manlu.org/tao/core/logx"
)

type CreateReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateReplyLogic {
	return &CreateReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateReplyLogic) CreateReply(in *reply.CreateReplyReq) (*reply.CreateReplyResp, error) {
	// todo: add your logic here and delete this line

	return &reply.CreateReplyResp{}, nil
}
