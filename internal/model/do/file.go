// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// File is the golang structure of table file for DAO operations like Where/Data.
type File struct {
	g.Meta    `orm:"table:file, do:true"`
	Id        interface{} // ID
	Name      interface{} // 图片名称
	Src       interface{} // 本地文件存储路径
	Url       interface{} // URL地址
	UserId    interface{} // 用户id
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
