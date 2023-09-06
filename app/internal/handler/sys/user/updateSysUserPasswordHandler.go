package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"laravel-single/app/internal/logic/sys/user"
	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"
)

func UpdateSysUserPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSysUserPasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateSysUserPasswordLogic(r.Context(), svcCtx)
		err := l.UpdateSysUserPassword(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
