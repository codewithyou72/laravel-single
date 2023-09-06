package job

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"laravel-single/app/internal/logic/sys/job"
	"laravel-single/app/internal/svc"
	"laravel-single/app/internal/types"
)

func GetSysJobPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysJobPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := job.NewGetSysJobPageLogic(r.Context(), svcCtx)
		resp, err := l.GetSysJobPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
