package goods

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model"
)

//type GoodLogic struct {
//	repo *dao.GoodsRepo
//}
//
//func New(repo *dao.GoodsRepo) *GoodLogic {
//	return &GoodLogic{
//		repo: repo,
//	}
//}
//
//func (l *GoodLogic) AddGoods(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error) {
//	lastInsertID, err := l.repo.Create(ctx, in)
//	if err != nil {
//		return out, err
//	}
//	return model.GoodsCreateOutput{Id: uint(lastInsertID)}, err
//}
//
//func (l *GoodLogic) GoodsDetail(ctx context.Context, in model.GoodsGetInput) (model.GoodsGetOutput, error) {
//	return l.repo.Get(ctx, in)
//}

type sGoods struct {
}

func New() *sGoods {
	return &sGoods{}
}

func (s *sGoods) AddGoods(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error) {
	lastInsertID, err := dao.GoodsRepo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.GoodsCreateOutput{Id: uint(lastInsertID)}, err
}

func (s *sGoods) GoodsDetail(ctx context.Context, in model.GoodsGetInput) (model.GoodsGetOutput, error) {
	var out model.GoodsGetOutput
	err := dao.GoodsRepo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	if err != nil {
		return model.GoodsGetOutput{}, err
	}
	return out, nil
}
