package logic

import (
	"context"

	"sllt/backend-learning/blog/rpc/article/article"
	"sllt/backend-learning/blog/rpc/article/internal/svc"

	"manlu.org/tao/core/logx"
)

type GetAllArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllArticleLogic {
	return &GetAllArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取全部文章
func (l *GetAllArticleLogic) GetAllArticle(in *article.GetAllArticleReq) (*article.GetAllArticleResp, error) {
	// todo: add your logic here and delete this line

	return &article.GetAllArticleResp{}, nil
}
