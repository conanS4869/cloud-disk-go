package handler

import (
	"cloud-disk-go/core/helper"
	"errors"
	"net/http"

	"cloud-disk-go/core/internal/logic"
	"cloud-disk-go/core/internal/svc"
	"cloud-disk-go/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadChunkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 参加必填校验
		if r.PostForm.Get("key") == "" {
			httpx.Error(w, errors.New("key is empty"))
		}
		if r.PostForm.Get("upload_id") == "" {
			httpx.Error(w, errors.New("upload_id is empty"))
		}
		if r.PostForm.Get("part_number") == "" {
			httpx.Error(w, errors.New("part_number is empty"))
		}

		etag, err := helper.CosPartUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadChunk(&req)
		resp = new(types.FileUploadChunkResponse)
		resp.Etag = etag

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
