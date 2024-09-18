package user

import (
	"net/http"

	"zero-demo/api/admin/internal/logic/system/user"
	"zero-demo/api/admin/internal/response"
	"zero-demo/api/admin/internal/svc"
	"zero-demo/api/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDeleteUserLogic(r.Context(), svcCtx)
		err := l.DeleteUser(&req)
		response.Response(w, nil, err) //②
	}
}
