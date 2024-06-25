package goods

import (
	"context"
	v1 "goframe_shop/api/goods/v1"
)

type IGoodsV1 interface {
	List(ctx context.Context, req *v1.GoodsGetListReq) (*v1.GoodsGetListRes, error)
	Update(ctx context.Context, req *v1.GoodsUpdateReq) (*v1.GoodsUpdateRes, error)
	Add(ctx context.Context, req *v1.GoodsAddReq) (*v1.GoodsAddRes, error)
	Get(ctx context.Context, req *v1.GoodsGetReq) (*v1.GoodsGetRes, error)
	Delete(ctx context.Context, req *v1.GoodsDeleteReq) (*v1.GoodsDeleteRes, error)
}
