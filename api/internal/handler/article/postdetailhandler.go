package article

import (
	"net/http"

	"manlu.org/tao/rest/httpx"
	"sllt/backend-learning/blog/api/internal/logic/article"
	"sllt/backend-learning/blog/api/internal/svc"
	"sllt/backend-learning/blog/api/internal/types"
)

func PostDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := article.NewPostDetailLogic(r.Context(), svcCtx)
		resp, err := l.PostDetail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
