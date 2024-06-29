package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Permission struct {
	g.Meta    `orm:"table:permission, do:true"`
	Id        interface{} //
	Name      interface{} // 权限名称
	Path      interface{} // 路径
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
