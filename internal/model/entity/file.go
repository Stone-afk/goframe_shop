// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// File is the golang structure for table file.
type File struct {
	Id        int         `json:"id"        orm:"id"         description:"ID"`
	Name      string      `json:"name"      orm:"name"       description:"图片名称"`
	Src       string      `json:"src"       orm:"src"        description:"本地文件存储路径"`
	Url       string      `json:"url"       orm:"url"        description:"URL地址"`
	UserId    int         `json:"userId"    orm:"user_id"    description:"用户id"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
