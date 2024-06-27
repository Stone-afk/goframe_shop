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
	IAdmin interface {
		Create(ctx context.Context, in model.AdminCreateInput) (model.AdminCreateOutput, error)
		GetUserByUserNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{}
		Delete(ctx context.Context, id uint) error
		Update(ctx context.Context, in model.AdminUpdateInput) error
		// List 查询内容列表
		List(ctx context.Context, in model.AdminGetListInput) (*model.AdminGetListOutput, error)
		GetAdminByNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{}
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
