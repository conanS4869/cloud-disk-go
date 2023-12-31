package handler

import (
	"net/http"

	"cloud-disk-go/core/internal/logic"
	"cloud-disk-go/core/internal/svc"
	"cloud-disk-go/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFolderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFolderListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFolderListLogic(r.Context(), svcCtx)
		resp, err := l.UserFolderList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
