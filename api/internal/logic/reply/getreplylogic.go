package reply

import (
	"context"

	"sllt/backend-learning/blog/api/internal/svc"
	"sllt/backend-learning/blog/api/internal/types"

	"manlu.org/tao/core/logx"
)

type GetReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetReplyLogic {
	return GetReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetReplyLogic) GetReply(req types.GetReplyReq) (resp *types.GetReplyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
