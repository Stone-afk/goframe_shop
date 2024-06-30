package user

import (
	"context"
	v1 "goframe_shop/api/frontend/user/v1"
)

type IUserV1 interface {
	Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterRes, error)
	Info(ctx context.Context, req *v1.UserInfoReq) (*v1.UserInfoRes, error)
	UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (*v1.UpdatePasswordRes, error)
}
