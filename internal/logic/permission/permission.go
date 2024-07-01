package permission

import (
	"context"
	"goframe_shop/internal/dao"
	"goframe_shop/internal/model"
	"goframe_shop/internal/model/entity"
	"goframe_shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

func (s *sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (model.PermissionCreateOutput, error) {
	lastInsertID, err := dao.PermissionRepo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.PermissionCreateOutput{}, err
	}
	return model.PermissionCreateOutput{PermissionId: uint(lastInsertID)}, err
}

func (s *sPermission) Delete(ctx context.Context, id uint) error {
	_, err := dao.PermissionRepo.Ctx(ctx).Where(g.Map{
		dao.PermissionRepo.Columns().Id: id,
	}).Delete()
	return err
}

func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	_, err := dao.PermissionRepo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.PermissionRepo.Columns().Id).
		Where(dao.PermissionRepo.Columns().Id, in.Id).
		Update()
	return err
}

func (s *sPermission) List(ctx context.Context, in model.PermissionGetListInput) (*model.PermissionGetListOutput, error) {
	var (
		m = dao.PermissionRepo.Ctx(ctx)
	)
	out := &model.PermissionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.Permission
	err := listModel.Scan(&list)
	if err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	//不指定item的键名用：Scan
	err = listModel.Scan(&out.List)
	if err != nil {
		return out, err
	}
	return out, nil
}
