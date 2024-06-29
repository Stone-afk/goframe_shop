package entity

import "github.com/gogf/gf/v2/os/gtime"

type RolePermission struct {
	Id           int         `json:"id"           description:""`
	RoleId       int         `json:"roleId"       description:"角色id"`
	PermissionId int         `json:"permissionId" description:"权限id"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`
}
