package goods

import (
	"context"
	"demo/api/goods"
	v1 "demo/api/goods/v1"
	"demo/internal/model"
	"demo/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type ControllerV1 struct{}

func (c ControllerV1) GoodsAddReq(ctx context.Context, req *v1.GoodsAddReq) (res *v1.GoodsAddRes, err error) {
	data := model.GoodsCreateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Goods().AddGoods(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.GoodsAddRes{Id: out.Id}, nil
}

func (c ControllerV1) GoodsGetReq(ctx context.Context, req *v1.GoodsGetReq) (res *v1.GoodsGetRes, err error) {
	detail, err := service.Goods().GoodsDetail(ctx, model.GoodsGetInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.GoodsGetRes{}
	err = gconv.Struct(detail, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewV1() goods.IGoodsV1 {
	return &ControllerV1{}
}
