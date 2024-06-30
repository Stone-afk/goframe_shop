package admin

import (
	"context"
	"goframe_shop/api/backend/admin"
	v1 "goframe_shop/api/backend/admin/v1"
	"goframe_shop/internal/consts"
	"goframe_shop/internal/model"
	"goframe_shop/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type ControllerV1 struct{}

func (c *ControllerV1) Get(ctx context.Context, req *v1.AdminGetReq) (*v1.AdminGetRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ControllerV1) List(ctx context.Context, req *v1.AdminGetListReq) (*v1.AdminGetListRes, error) {
	getListRes, err := service.Admin().List(ctx, model.AdminGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &v1.AdminGetListRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.AdminUpdateReq) (*v1.AdminUpdateRes, error) {
	err := service.Admin().Update(ctx, model.AdminUpdateInput{
		Id: req.Id,
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	return &v1.AdminUpdateRes{Id: req.Id}, err
}

func (c *ControllerV1) Add(ctx context.Context, req *v1.AdminAddReq) (*v1.AdminAddRes, error) {
	out, err := service.Admin().Create(ctx, model.AdminCreateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.AdminAddRes{AdminId: out.AdminId}, nil
}

func (c *ControllerV1) Info(ctx context.Context, req *v1.AdminGetReq) (*v1.AdminGetRes, error) {
	return &v1.AdminGetRes{
		Id:      gconv.Int(ctx.Value(consts.CtxAdminId)),
		Name:    gconv.String(ctx.Value(consts.CtxAdminName)),
		IsAdmin: gconv.Int(ctx.Value(consts.CtxAdminIsAdmin)),
		RoleIds: gconv.String(ctx.Value(consts.CtxAdminRoleIds)),
	}, nil
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.AdminDeleteReq) (*v1.AdminDeleteRes, error) {
	return nil, service.Admin().Delete(ctx, req.Id)
}

func NewV1() admin.IAdminV1 {
	return &ControllerV1{}
}
