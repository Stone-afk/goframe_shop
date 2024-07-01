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
	IPermission interface {
		Create(ctx context.Context, in model.PermissionCreateInput) (model.PermissionCreateOutput, error)
		Delete(ctx context.Context, id uint) error
		Update(ctx context.Context, in model.PermissionUpdateInput) error
		List(ctx context.Context, in model.PermissionGetListInput) (*model.PermissionGetListOutput, error)
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
