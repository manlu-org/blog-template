package article

import (
	"context"

	"sllt/backend-learning/blog/api/internal/svc"
	"sllt/backend-learning/blog/api/internal/types"

	"manlu.org/tao/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreatePostLogic {
	return CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req types.CreatePostReq) (resp *types.CreatePostResp, err error) {
	// todo: add your logic here and delete this line
	return
}
