package file

import (
	"context"
	v1 "goframe_shop/api/backend/file/v1"
)

type IFileV1 interface {
	Upload(ctx context.Context, req *v1.FileUploadReq) (*v1.FileUploadRes, error)
}
