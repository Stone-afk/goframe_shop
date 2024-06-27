package dao

import "goframe_shop/internal/dao/internal"

// internalUserInfoDao is internal type for wrapping internal DAO implements.
type internalUserDao = *internal.UserDao

// userInfoDao is the data access object for table user_info.
// You can define custom methods on it to extend its functionality as you wish.
type userRepo struct {
	internalUserDao
}

var (
	UserRepo = userRepo{
		internal.NewUserDao(),
	}
)
