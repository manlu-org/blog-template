package logic

import (
	"context"
	"fmt"

	"sllt/backend-learning/blog/rpc/reply/internal/svc"
	"sllt/backend-learning/blog/rpc/reply/reply"

	"manlu.org/tao/core/logx"
)

type GetRepliesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRepliesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRepliesLogic {
	return &GetRepliesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRepliesLogic) GetReplies(in *reply.GetRepliesReq) (*reply.GetRepliesResp, error) {

	replies := make([]*reply.ReplyInfo, 0)
	for i := 1; i < 3; i++ {
		reply := &reply.ReplyInfo{
			Id:      int64(i),
			Content: fmt.Sprintf("hello%d", i),
		}
		replies = append(replies, reply)
	}

	resp := &reply.GetRepliesResp{
		List: replies,
	}

	return resp, nil
}
