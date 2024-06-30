package login

import (
	"context"
	"goframe_shop/api/login"
	backend "goframe_shop/api/login/v1"
	"goframe_shop/internal/service"
)

type ControllerV1 struct{}

func (c *ControllerV1) Do(ctx context.Context, req *backend.LoginDoReq) (*backend.LoginDoRes, error) {
	//TODO implement me
	panic("implement me")
}

// for session
//func (a *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
//	res = &backend.LoginDoRes{}
//	err = service.Login().Login(ctx, model.UserLoginInput{
//		Name:     req.Name,
//		Password: req.Password,
//	})
//	if err != nil {
//		return
//	}
//	// 识别并跳转到登录前页面
//	//res.Info = service.Session().GetUser(ctx)
//	return
//}

//for jwt
//func (c *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
//	res = &backend.LoginDoRes{}
//	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
//	return
//}

func (c *ControllerV1) RefreshToken(ctx context.Context, req *backend.RefreshTokenReq) (res *backend.RefreshTokenRes, err error) {
	res = &backend.RefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

func (c *ControllerV1) Logout(ctx context.Context, req *backend.LogoutReq) (res *backend.LogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}

func NewV1() login.ILoginV1 {
	return &ControllerV1{}
}
