package role

import (
	"context"
	"goframe_shop/api/backend/role"
	v1 "goframe_shop/api/backend/role/v1"
	"goframe_shop/internal/model"
	"goframe_shop/internal/service"
)

type ControllerV1 struct{}

func (c *ControllerV1) List(ctx context.Context, req *v1.RoleGetListCommonReq) (*v1.RoleGetListCommonRes, error) {
	getListRes, err := service.Role().List(ctx, model.RoleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &v1.RoleGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.RoleUpdateReq) (*v1.RoleUpdateRes, error) {
	err := service.Role().Update(ctx, model.RoleUpdateInput{
		Id: req.Id,
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	return &v1.RoleUpdateRes{Id: req.Id}, err
}

func (c *ControllerV1) Add(ctx context.Context, req *v1.RoleReq) (*v1.RoleRes, error) {
	out, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.RoleRes{RoleId: out.RoleId}, nil
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.RoleDeleteReq) (*v1.RoleDeleteRes, error) {
	err := service.Role().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.RoleDeleteRes{}, nil
}

func (c *ControllerV1) AddPermission(ctx context.Context, req *v1.AddPermissionReq) (*v1.AddPermissionRes, error) {
	permission, err := service.Role().AddPermission(ctx, model.RoleAddPermissionInput{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AddPermissionRes{Id: permission.Id}, err
}

func (c *ControllerV1) DeletePermission(ctx context.Context, req *v1.DeletePermissionReq) (*v1.DeletePermissionRes, error) {
	err := service.Role().DeletePermission(ctx, model.RoleDeletePermissionInput{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeletePermissionRes{}, err
}

func NewV1() role.IRoleV1 {
	return &ControllerV1{}
}
