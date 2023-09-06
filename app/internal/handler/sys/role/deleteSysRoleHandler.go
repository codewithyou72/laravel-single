package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"laravel-single/app/internal/logic/sys/role"
	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"
)

func DeleteSysRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteSysRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewDeleteSysRoleLogic(r.Context(), svcCtx)
		err := l.DeleteSysRole(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
