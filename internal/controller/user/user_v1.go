package user

import (
	"context"
	"goframe_shop/api/frontend/user"
	v1 "goframe_shop/api/frontend/user/v1"
	"goframe_shop/internal/consts"
	"goframe_shop/internal/model"
	"goframe_shop/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type ControllerV1 struct{}

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterRes, error) {
	in := model.RegisterInput{}
	err := gconv.Struct(req, &in)
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterRes{Id: out.Id}, nil
}

func (c *ControllerV1) Info(ctx context.Context, req *v1.UserInfoReq) (*v1.UserInfoRes, error) {
	res := &v1.UserInfoRes{}
	res.Id = gconv.Uint(ctx.Value(consts.CtxUserId))
	res.Name = gconv.String(ctx.Value(consts.CtxUserName))
	res.Avatar = gconv.String(ctx.Value(consts.CtxUserAvatar))
	res.Sex = gconv.Uint8(ctx.Value(consts.CtxUserSex))
	res.Sign = gconv.String(ctx.Value(consts.CtxUserSign))
	res.Status = gconv.Uint8(ctx.Value(consts.CtxUserStatus))
	return res, nil
}

func (c *ControllerV1) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (*v1.UpdatePasswordRes, error) {
	in := model.UpdatePasswordInput{}
	err := gconv.Struct(req, &in)
	if err != nil {
		return nil, err
	}
	out, err := service.User().UpdatePassword(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.UpdatePasswordRes{Id: out.Id}, nil
}

func NewV1() user.IUserV1 {
	return &ControllerV1{}
}
