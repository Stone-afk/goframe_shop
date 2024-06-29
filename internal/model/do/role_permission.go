package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type RolePermission struct {
	g.Meta       `orm:"table:role_permission, do:true"`
	Id           interface{} //
	RoleId       interface{} // 角色id
	PermissionId interface{} // 权限id
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
