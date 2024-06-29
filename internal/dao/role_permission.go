package dao

import "goframe_shop/internal/dao/internal"

// internalRolePermissionInfoDao is internal type for wrapping internal DAO implements.
type internalRolePermissionDao = *internal.RolePermissionDao

// rolePermissionInfoDao is the data access object for table role_permission_info.
// You can define custom methods on it to extend its functionality as you wish.
type rolePermissionRepo struct {
	internalRolePermissionDao
}

var (
	RolePermissionRepo = rolePermissionRepo{
		internal.NewRolePermissionDao(),
	}
)
