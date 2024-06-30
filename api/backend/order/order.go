package order

import (
	"context"
	v1 "goframe_shop/api/backend/order/v1"
)

type IOrderV1 interface {
	List(ctx context.Context, req *v1.OrderListReq) (*v1.OrderListRes, error)
	Get(ctx context.Context, req *v1.OrderDetailReq) (*v1.OrderDetailRes, error)
}
