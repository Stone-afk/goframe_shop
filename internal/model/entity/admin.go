package entity

import "github.com/gogf/gf/v2/os/gtime"

type Admin struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:"用户名"`
	Password  string      `json:"password"  description:"密码"`
	RoleIds   string      `json:"roleIds"   description:"角色ids"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	UserSalt  string      `json:"userSalt"  description:"加密盐"`
	IsAdmin   int         `json:"isAdmin"   description:"是否超级管理员"`
}
