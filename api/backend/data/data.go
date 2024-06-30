package data

import (
	"context"
	v1 "goframe_shop/api/backend/data/v1"
)

type IDataV1 interface {
	DataHead(ctx context.Context, req *v1.DataHeadReq) (*v1.DataHeadRes, error)
	Echarts(ctx context.Context, req *v1.DataEChartsReq) (*v1.DataEChartsRes, error)
}
