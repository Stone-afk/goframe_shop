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
	IRole interface {
		Create(ctx context.Context, in model.RoleCreateInput) (model.RoleCreateOutput, error)
		Delete(ctx context.Context, id uint) error
		Update(ctx context.Context, in model.RoleUpdateInput) error
		// List 查询内容列表
		List(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error)
		AddPermission(ctx context.Context, in model.RoleAddPermissionInput) (model.RoleAddPermissionOutput, error)
		DeletePermission(ctx context.Context, in model.RoleDeletePermissionInput) error
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
