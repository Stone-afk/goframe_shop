package permission

import (
	"context"
	"goframe_shop/api/backend/permission"
	v1 "goframe_shop/api/backend/permission/v1"
	"goframe_shop/internal/model"
	"goframe_shop/internal/service"
)

type ControllerV1 struct{}

func (c *ControllerV1) List(ctx context.Context, req *v1.PermissionGetListCommonReq) (*v1.PermissionGetListCommonRes, error) {
	getListRes, err := service.Permission().List(ctx, model.PermissionGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &v1.PermissionGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.PermissionUpdateReq) (*v1.PermissionUpdateRes, error) {
	err := service.Permission().Update(ctx, model.PermissionUpdateInput{
		Id: req.Id,
		PermissionCreateUpdateBase: model.PermissionCreateUpdateBase{
			Name: req.Name,
			Path: req.Path,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.PermissionUpdateRes{Id: req.Id}, nil
}

func (c *ControllerV1) Add(ctx context.Context, req *v1.PermissionReq) (*v1.PermissionRes, error) {
	out, err := service.Permission().Create(ctx, model.PermissionCreateInput{
		PermissionCreateUpdateBase: model.PermissionCreateUpdateBase{
			Name: req.Name,
			Path: req.Path,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.PermissionRes{PermissionId: out.PermissionId}, nil
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.PermissionDeleteReq) (*v1.PermissionDeleteRes, error) {
	err := service.Permission().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.PermissionDeleteRes{}, nil
}

func NewV1() permission.IPermissionV1 {
	return &ControllerV1{}
}
