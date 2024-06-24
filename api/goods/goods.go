package goods

import (
	"context"
	v1 "goframe_shop/api/goods/v1"
)

type IGoodsV1 interface {
	GoodsAddReq(ctx context.Context, req *v1.GoodsAddReq) (*v1.GoodsAddRes, error)
	GoodsGetReq(ctx context.Context, req *v1.GoodsGetReq) (*v1.GoodsGetRes, error)
	GoodsDeleteReq(ctx context.Context, req *v1.GoodsDeleteReq) (*v1.GoodsDeleteRes, error)
}
