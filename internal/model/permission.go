package model

// PermissionCreateUpdateBase 创建/修改内容基类
type PermissionCreateUpdateBase struct {
	Name string
	Path string
}

// PermissionCreateInput 创建内容
type PermissionCreateInput struct {
	PermissionCreateUpdateBase
}

// PermissionCreateOutput 创建内容返回结果
type PermissionCreateOutput struct {
	PermissionId uint `json:"permission_id"`
}

// PermissionUpdateInput 修改内容
type PermissionUpdateInput struct {
	PermissionCreateUpdateBase
	Id uint
}
