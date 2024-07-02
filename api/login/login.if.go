// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package login

import (
	"context"

	"goframe_shop/api/login/v1"
)

type ILoginV1 interface {
	LoginDo(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error)
	RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
}
