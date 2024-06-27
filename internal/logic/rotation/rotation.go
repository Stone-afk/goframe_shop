package rotation

import (
	"context"
	"goframe_shop/internal/dao"
	"goframe_shop/internal/model"
	"goframe_shop/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
)

type sRotation struct{}

func New() *sRotation {
	return &sRotation{}
}

func init() {
	service.RegisterRotation(New())
}

func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (model.RotationCreateOutput, error) {
	var out model.RotationCreateOutput
	if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationRepo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: int(lastInsertID)}, err
}

func (s *sRotation) Delete(ctx context.Context, id uint) error {
	return dao.RotationRepo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.RotationRepo.Ctx(ctx).Where(g.Map{
			dao.RotationRepo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationRepo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.RotationRepo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationRepo.Columns().Id).
			Where(dao.RotationRepo.Columns().Id, in.Id).
			Update()
		return err
	})
}

func (s *sRotation) List(ctx context.Context, in model.RotationGetListInput) (*model.RotationGetListOutput, error) {
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.RotationRepo.Ctx(ctx)
	out := &model.RotationGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	listModel := m.Page(in.Page, in.Size)
	var err error
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	out.List = make([]model.RotationGetListOutputItem, 0, in.Size)
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return out, nil
}
