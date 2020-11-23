package handler

import (
	"net/http"

	"github.com/reaperhero/go-shorturl/api/internal/logic"
	"github.com/reaperhero/go-shorturl/api/internal/svc"
	"github.com/reaperhero/go-shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func expandHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewExpandLogic(r.Context(), ctx)
		resp, err := l.Expand(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
