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

// List 查询分类列表
func (s *sGoods) List(ctx context.Context, in model.GoodsGetListInput) (*model.GoodsGetListOutput, error) {
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.GoodsRepo.Ctx(ctx)
	//2. 实例化响应结构体
	out := &model.GoodsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//3. 分页查询
	listModel := m.Page(in.Page, in.Size)
	//4. 再查询count，判断有无数据
	var err error
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	//5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.GoodsGetListOutputItem, 0, in.Size)
	//6. 把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return out, nil
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
