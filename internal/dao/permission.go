package dao

import "goframe_shop/internal/dao/internal"

type internalPermissionDao = *internal.PermissionDao

// permissionInfoDao is the data access object for table permission_info.
// You can define custom methods on it to extend its functionality as you wish.
type permissionRepo struct {
	internalPermissionDao
}

var (
	PermissionRepo = permissionRepo{
		internal.NewPermissionDao(),
	}
)
