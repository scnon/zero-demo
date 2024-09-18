package user

import (
	"net/http"

	"zero-demo/api/admin/internal/logic/system/user"
	"zero-demo/api/admin/internal/response"
	"zero-demo/api/admin/internal/svc"
	"zero-demo/api/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryUserListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewQueryUserListLogic(r.Context(), svcCtx)
		resp, err := l.QueryUserList(&req)
		response.Response(w, resp, err) //②
	}
}
