package role

import (
	"context"
	"goframe_shop/internal/dao"
	"goframe_shop/internal/model"
	"goframe_shop/internal/model/entity"
	"goframe_shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (model.RoleCreateOutput, error) {
	lastInsertID, err := dao.RoleRepo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.RoleCreateOutput{}, err
	}
	return model.RoleCreateOutput{RoleId: uint(lastInsertID)}, err
}

func (s *sRole) Delete(ctx context.Context, id uint) error {
	// 删除内容
	_, err := dao.RoleRepo.Ctx(ctx).Where(g.Map{
		dao.RoleRepo.Columns().Id: id,
	}).Unscoped().Delete()
	return err
}

// GetList 查询内容列表
func (s *sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	var (
		m = dao.RoleRepo.Ctx(ctx)
	)
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.Role
	if err := listModel.Scan(&list); err != nil {
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
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (s *sRole) AddPermission(ctx context.Context, in model.RoleAddPermissionInput) (model.RoleAddPermissionOutput, error) {
	id, err := dao.RolePermissionRepo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.RoleAddPermissionOutput{}, err
	}
	return model.RoleAddPermissionOutput{Id: uint(id)}, err
}

func (s *sRole) DeletePermission(ctx context.Context, in model.RoleDeletePermissionInput) error {
	_, err := dao.RolePermissionRepo.Ctx(ctx).Where(g.Map{
		dao.RolePermissionRepo.Columns().RoleId:       in.RoleId,
		dao.RolePermissionRepo.Columns().PermissionId: in.PermissionId,
	}).Delete()
	if err != nil {
		return err
	}
	return nil
}
