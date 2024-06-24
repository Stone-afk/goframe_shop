package dao

import (
	"goframe_shop/internal/dao/internal"
)

// internalGoodsInfoDao is internal type for wrapping internal DAO implements.
type internalGoodsDao = *internal.GoodsDao

type goodsRepo struct {
	internalGoodsDao
}

var (
	GoodsRepo = goodsRepo{
		internal.NewGoodsDao(),
	}
)
