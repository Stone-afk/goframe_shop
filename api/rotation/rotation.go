package v1

import (
	"context"
	v1 "goframe_shop/api/rotation/v1"
)

type IRotationV1 interface {
	List(ctx context.Context, req *v1.RotationGetListReq) (*v1.RotationGetListRes, error)
	Update(ctx context.Context, req *v1.RotationUpdateReq) (*v1.RotationUpdateRes, error)
	Add(ctx context.Context, req *v1.RotationAddReq) (*v1.RotationAddRes, error)
	Delete(ctx context.Context, req *v1.RotationDeleteReq) (*v1.RotationDeleteRes, error)
}
