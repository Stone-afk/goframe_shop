package dao

import "goframe_shop/internal/dao/internal"

// internalRotationInfoDao is internal type for wrapping internal DAO implements.
type internalRotationDao = *internal.RotationDao

// rotationInfoDao is the data access object for table rotation_info.
// You can define custom methods on it to extend its functionality as you wish.
type rotationInfoRepo struct {
	internalRotationDao
}

var (
	RotationRepo = rotationInfoRepo{
		internal.NewRotationDao(),
	}
)
