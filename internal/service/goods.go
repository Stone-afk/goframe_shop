// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"demo/internal/model"
)

type (
	IGoods interface {
		AddGoods(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error)
		GoodsDetail(ctx context.Context, in model.GoodsGetInput) (model.GoodsGetOutput, error)
	}
)

var (
	localGoods IGoods
)

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
