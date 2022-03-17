package reply

import (
	"context"

	"sllt/backend-learning/blog/api/internal/svc"
	"sllt/backend-learning/blog/api/internal/types"

	"manlu.org/tao/core/logx"
)

type CreateReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateReplyLogic {
	return CreateReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateReplyLogic) CreateReply(req types.CreateReplyReq) (resp *types.CreateReplyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
