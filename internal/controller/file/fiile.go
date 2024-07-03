package file

import (
	"context"
	"goframe_shop/api/backend/file"
	v1 "goframe_shop/api/backend/file/v1"
	"goframe_shop/internal/model"
	"goframe_shop/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type ControllerV1 struct{}

func (c *ControllerV1) Upload(ctx context.Context, req *v1.FileUploadReq) (*v1.FileUploadRes, error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
	}
	upload, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	return &v1.FileUploadRes{
		Name: upload.Name,
		Url:  upload.Url,
	}, nil
}

func NewV1() file.IFileV1 {
	return &ControllerV1{}
}
