package permission

import (
	"context"
	v1 "goframe_shop/api/backend/permission/v1"
)

type IPermissionV1 interface {
	List(ctx context.Context, req *v1.PermissionGetListCommonReq) (*v1.PermissionGetListCommonRes, error)
	Update(ctx context.Context, req *v1.PermissionUpdateReq) (*v1.PermissionUpdateRes, error)
	Add(ctx context.Context, req *v1.PermissionReq) (*v1.PermissionRes, error)
	Delete(ctx context.Context, req *v1.PermissionDeleteReq) (*v1.PermissionDeleteRes, error)
}
