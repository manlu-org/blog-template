package reply

import (
	"net/http"

	"manlu.org/tao/rest/httpx"
	"sllt/backend-learning/blog/api/internal/logic/reply"
	"sllt/backend-learning/blog/api/internal/svc"
	"sllt/backend-learning/blog/api/internal/types"
)

func GetReplyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReplyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := reply.NewGetReplyLogic(r.Context(), svcCtx)
		resp, err := l.GetReply(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
