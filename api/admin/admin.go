package admin

import (
	"context"
	v1 "goframe_shop/api/Admin/v1"
)

type IAdminV1 interface {
	List(ctx context.Context, req *v1.AdminGetListReq) (*v1.AdminGetListRes, error)
	Update(ctx context.Context, req *v1.AdminUpdateReq) (*v1.AdminUpdateRes, error)
	Add(ctx context.Context, req *v1.AdminAddReq) (*v1.AdminAddRes, error)
	Get(ctx context.Context, req *v1.AdminGetReq) (*v1.AdminGetRes, error)
	Delete(ctx context.Context, req *v1.AdminDeleteReq) (*v1.AdminDeleteRes, error)
}
