package rotation

import (
	"context"
	"goframe_shop/api/backend/rotation"
	v1 "goframe_shop/api/backend/rotation/v1"
	"goframe_shop/internal/model"
	"goframe_shop/internal/service"
)

type ControllerV1 struct{}

func (c *ControllerV1) List(ctx context.Context, req *v1.RotationGetListReq) (*v1.RotationGetListRes, error) {
	getListRes, err := service.Rotation().List(ctx, model.RotationGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RotationGetListRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.RotationUpdateReq) (*v1.RotationUpdateRes, error) {
	err := service.Rotation().Update(ctx, model.RotationUpdateInput{
		Id: req.Id,
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	return &v1.RotationUpdateRes{Id: req.Id}, err
}

func (c *ControllerV1) Add(ctx context.Context, req *v1.RotationAddReq) (*v1.RotationAddRes, error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.RotationAddRes{RotationId: out.RotationId}, nil
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.RotationDeleteReq) (*v1.RotationDeleteRes, error) {
	return nil, service.Rotation().Delete(ctx, req.Id)
}

func NewV1() rotation.IRotationV1 {
	return &ControllerV1{}
}
