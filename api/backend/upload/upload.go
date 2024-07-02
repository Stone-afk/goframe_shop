package upload

import (
	"context"
	v1 "goframe_shop/api/backend/upload/v1"
)

type IUploadV1 interface {
	UploadImgToCloud(ctx context.Context, req *v1.UploadImgToCloudReq) (*v1.UploadImgToCloudRes, error)
}
