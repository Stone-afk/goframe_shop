package role

import (
	"context"
	v1 "goframe_shop/api/backend/role/v1"
)

type IRoleV1 interface {
	List(ctx context.Context, req *v1.RoleGetListCommonReq) (*v1.RoleGetListCommonRes, error)
	Update(ctx context.Context, req *v1.RoleUpdateReq) (*v1.RoleUpdateRes, error)
	Add(ctx context.Context, req *v1.RoleReq) (*v1.RoleRes, error)
	Delete(ctx context.Context, req *v1.RoleDeleteReq) (*v1.RoleDeleteRes, error)
	AddPermission(ctx context.Context, req *v1.AddPermissionReq) (*v1.AddPermissionRes, error)
	DeletePermission(ctx context.Context, req *v1.DeletePermissionReq) (*v1.DeletePermissionRes, error)
}
