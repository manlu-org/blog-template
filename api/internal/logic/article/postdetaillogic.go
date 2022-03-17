package article

import (
	"context"
	"sllt/backend-learning/blog/rpc/article/articleclient"

	"sllt/backend-learning/blog/api/internal/svc"
	"sllt/backend-learning/blog/api/internal/types"

	"manlu.org/tao/core/logx"
)

type PostDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostDetailLogic {
	return PostDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostDetailLogic) PostDetail(req types.PostDetailReq) (resp *types.PostDetailResp, err error) {

	r := &articleclient.GetArticleReq{
		Id: int32(req.Id),
	}
	result, err := l.svcCtx.Article.GetArticle(l.ctx, r)
	if err != nil {
		return nil, err
	}

	replies := make([]*types.Reply, 0)

	for _, res := range result.Post.Replies {
		reply := &types.Reply{
			Id:      int64(res.Id),
			Content: res.Content,
		}
		replies = append(replies, reply)
	}

	resp = &types.PostDetailResp{
		Id:        int64(result.Post.Id),
		Title:     result.Post.Title,
		Content:   result.Post.Content,
		ReplyList: replies,
	}
	return
}
