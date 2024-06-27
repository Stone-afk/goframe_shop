package login

import (
	"context"
	v1 "goframe_shop/api/login/v1"
)

type ILoginV1 interface {
	Do(ctx context.Context, req *v1.LoginDoReq) (*v1.LoginDoRes, error)
	RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (*v1.RefreshTokenRes, error)
	Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutRes, error)
}
