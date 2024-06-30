package order

import (
	"context"
	v1 "goframe_shop/api/frontend/order/v1"
)

type IOrderV1 interface {
	Add(ctx context.Context, req *v1.AddOrderReq) (*v1.AddOrderRes, error)
}
