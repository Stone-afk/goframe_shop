package admin

import (
	"context"
	v12 "goframe_shop/api/backend/admin/v1"
)

type IAdminV1 interface {
	List(ctx context.Context, req *v12.AdminGetListReq) (*v12.AdminGetListRes, error)
	Update(ctx context.Context, req *v12.AdminUpdateReq) (*v12.AdminUpdateRes, error)
	Add(ctx context.Context, req *v12.AdminAddReq) (*v12.AdminAddRes, error)
	Get(ctx context.Context, req *v12.AdminGetReq) (*v12.AdminGetRes, error)
	Delete(ctx context.Context, req *v12.AdminDeleteReq) (*v12.AdminDeleteRes, error)
}
