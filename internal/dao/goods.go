package dao

import (
	"demo/internal/dao/internal"
)

//type GoodsRepo struct {
//	dao *internal.GoodsDao
//}
//
//func (repo *GoodsRepo) Create(ctx context.Context, in model.GoodsCreateInput) (int64, error) {
//	return repo.dao.Ctx(ctx).Data(in).InsertAndGetId()
//}
//
//func (repo *GoodsRepo) Get(ctx context.Context, in model.GoodsGetInput) (model.GoodsGetOutput, error) {
//	var out model.GoodsGetOutput
//	err := repo.dao.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
//	if err != nil {
//		return out, err
//	}
//	return out, nil
//}
//
//func NewGoodsRepo(dao *internal.GoodsDao) *GoodsRepo {
//	return &GoodsRepo{
//		dao: dao,
//	}
//}

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
