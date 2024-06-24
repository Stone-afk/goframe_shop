// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe_shop/internal/model"
)

type (
	IGoods interface {
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		Create(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error)
		Get(ctx context.Context, in model.GoodsGetInput) (model.GoodsGetOutput, error)
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
