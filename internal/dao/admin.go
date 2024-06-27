package dao

import "goframe_shop/internal/dao/internal"

// internalAdminInfoDao is internal type for wrapping internal DAO implements.
type internalAdminDao = *internal.AdminDao

// adminInfoDao is the data access object for table admin_info.
// You can define custom methods on it to extend its functionality as you wish.
type adminRepo struct {
	internalAdminDao
}

var (
	AdminRepo = adminRepo{
		internal.NewAdminDao(),
	}
)
