// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermission is the golang structure for table role_permission.
type RolePermission struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	RoleId       int         `json:"roleId"       orm:"role_id"       description:"角色id"`
	PermissionId int         `json:"permissionId" orm:"permission_id" description:"权限id"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`
}
