package logic

import (
	"context"
	"sllt/backend-learning/blog/rpc/article/articleclient"
	"sllt/backend-learning/blog/rpc/reply/replyclient"

	"sllt/backend-learning/blog/rpc/article/article"
	"sllt/backend-learning/blog/rpc/article/internal/svc"

	"manlu.org/tao/core/logx"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取文章详情
func (l *GetArticleLogic) GetArticle(in *article.GetArticleReq) (*article.GetArticleResp, error) {
	r := &replyclient.GetRepliesReq{
		PostId: int64(in.Id),
	}

	results, err := l.svcCtx.Reply.GetReplies(l.ctx, r)
	if err != nil {
		return nil, err
	}

	logx.Info(results)

	replies := make([]*articleclient.Reply, 0)
	for _, res := range results.List {
		reply := &articleclient.Reply{
			Id:      int32(res.Id),
			Content: res.Content,
		}
		replies = append(replies, reply)
	}

	a := &article.GetArticleResp{
		Post: &article.Post{
			Id:      in.Id,
			Title:   "hello",
			Content: "world",
			Replies: replies,
		}}
	return a, nil
}
