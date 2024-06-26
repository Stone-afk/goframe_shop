package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RotationInfo is the golang structure of table rotation_info for DAO operations like Where/Data.
type RotationInfo struct {
	g.Meta    `orm:"table:rotation_info, do:true"`
	Id        interface{} //
	PicUrl    interface{} // 轮播图片
	Link      interface{} // 跳转链接
	Sort      interface{} // 排序字段
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
