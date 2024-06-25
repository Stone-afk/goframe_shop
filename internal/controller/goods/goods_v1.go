package goods

import (
	"context"
	"goframe_shop/api/goods"
	v1 "goframe_shop/api/goods/v1"
	"goframe_shop/internal/model"
	"goframe_shop/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type ControllerV1 struct{}

func (c *ControllerV1) List(ctx context.Context, req *v1.GoodsGetListReq) (*v1.GoodsGetListRes, error) {
	getListRes, err := service.Goods().List(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GoodsGetListRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.GoodsUpdateReq) (*v1.GoodsUpdateRes, error) {
	in := model.GoodsUpdateInput{}
	err := gconv.Struct(req, &in)
	if err != nil {
		return nil, err
	}
	err = service.Goods().Update(ctx, in)
	return &v1.GoodsUpdateRes{Id: req.Id}, nil
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.GoodsDeleteReq) (*v1.GoodsDeleteRes, error) {
	err := service.Goods().Delete(ctx, req.Id)
	return &v1.GoodsDeleteRes{}, err
}

func (c *ControllerV1) Add(ctx context.Context, req *v1.GoodsAddReq) (*v1.GoodsAddRes, error) {
	in := model.GoodsCreateInput{}
	err := gconv.Scan(req, &in)
	if err != nil {
		return nil, err
	}
	out, err := service.Goods().Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return &v1.GoodsAddRes{Id: out.Id}, nil
}

func (c *ControllerV1) Get(ctx context.Context, req *v1.GoodsGetReq) (*v1.GoodsGetRes, error) {
	detail, err := service.Goods().Get(ctx, model.GoodsGetInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.GoodsGetRes{}
	err = gconv.Struct(detail, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewV1() goods.IGoodsV1 {
	return &ControllerV1{}
}
