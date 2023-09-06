package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"laravel-single/app/internal/logic/sys/menu"
	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"
)

func AddSysPermMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddSysPermMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewAddSysPermMenuLogic(r.Context(), svcCtx)
		err := l.AddSysPermMenu(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
