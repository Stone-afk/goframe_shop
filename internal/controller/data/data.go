package data

import (
	"context"
	"goframe_shop/api/backend/data"
	v1 "goframe_shop/api/backend/data/v1"
	"goframe_shop/internal/service"
)

type ControllerV1 struct{}

func (c *ControllerV1) Echarts(ctx context.Context, req *v1.DataEChartsReq) (*v1.DataEChartsRes, error) {
	echats, err := service.Data().Echarts(ctx)
	if err != nil {
		return &v1.DataEChartsRes{}, err
	}
	return &v1.DataEChartsRes{
		OrderTotal:           echats.OrderTotal,
		SalePriceTotal:       echats.SalePriceTotal,
		ConsumptionPerPerson: echats.ConsumptionPerPerson,
		NewOrder:             echats.NewOrder,
	}, err
}

func (c *ControllerV1) DataHead(ctx context.Context, req *v1.DataHeadReq) (res *v1.DataHeadRes, err error) {
	card, err := service.Data().HeadCard(ctx)
	if err != nil {
		return &v1.DataHeadRes{}, err
	}
	return &v1.DataHeadRes{
		TodayOrderCount: card.TodayOrderCount,
		DAU:             card.DAU,
		ConversionRate:  card.ConversionRate,
	}, err
}

func NewV1() data.IDataV1 {
	return &ControllerV1{}
}
