package goods

import (
	"context"
	"goframe_shop/internal/dao"
	"goframe_shop/internal/model"
	"goframe_shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterGoods(New())
}

type sGoods struct {
}

func New() *sGoods {
	return &sGoods{}
}

// Update 修改
func (s *sGoods) Update(ctx context.Context, in model.GoodsUpdateInput) error {
	_, err := dao.GoodsRepo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.GoodsRepo.Columns().Id).
		Where(dao.GoodsRepo.Columns().Id, in.Id).
		Update()
	return err
}

// Delete 删除
func (s *sGoods) Delete(ctx context.Context, id uint) error {
	_, err := dao.GoodsRepo.Ctx(ctx).Where(g.Map{
		dao.GoodsRepo.Columns().Id: id,
	}).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *sGoods) Create(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error) {
	lastInsertID, err := dao.GoodsRepo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.GoodsCreateOutput{Id: uint(lastInsertID)}, err
}

func (s *sGoods) Get(ctx context.Context, in model.GoodsGetInput) (model.GoodsGetOutput, error) {
	var out model.GoodsGetOutput
	err := dao.GoodsRepo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	if err != nil {
		return model.GoodsGetOutput{}, err
	}
	return out, nil
}
