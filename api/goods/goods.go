package goods

import (
	"context"
	v1 "demo/api/goods/v1"
)

type IGoodsV1 interface {
	GoodsAddReq(ctx context.Context, req *v1.GoodsAddReq) (res *v1.GoodsAddRes, err error)
	GoodsGetReq(ctx context.Context, req *v1.GoodsGetReq) (res *v1.GoodsGetRes, err error)
}
