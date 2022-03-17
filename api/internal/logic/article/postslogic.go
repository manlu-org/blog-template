package article

import (
	"context"

	"sllt/backend-learning/blog/api/internal/svc"
	"sllt/backend-learning/blog/api/internal/types"

	"manlu.org/tao/core/logx"
)

type PostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostsLogic {
	return PostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsLogic) Posts(req types.PostListReq) (resp *types.PostListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
