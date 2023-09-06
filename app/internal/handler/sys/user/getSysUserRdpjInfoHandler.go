package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"laravel-single/app/internal/logic/sys/user"
	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"
)

func GetSysUserRdpjInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSysUserRdpjInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetSysUserRdpjInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetSysUserRdpjInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
