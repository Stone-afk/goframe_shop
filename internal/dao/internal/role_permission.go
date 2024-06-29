package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type RolePermissionDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns RolePermissionColumns // columns contains all the column names of Table for convenient usage.
}

type RolePermissionColumns struct {
	Id           string //
	RoleId       string // 角色id
	PermissionId string // 权限id
	CreatedAt    string //
	UpdatedAt    string //
}

var rolePermissionColumns = RolePermissionColumns{
	Id:           "id",
	RoleId:       "role_id",
	PermissionId: "permission_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

func NewRolePermissionDao() *RolePermissionDao {
	return &RolePermissionDao{
		group:   "default",
		table:   "role_permission_info",
		columns: rolePermissionColumns,
	}
}

func (dao *RolePermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *RolePermissionDao) Table() string {
	return dao.table
}

func (dao *RolePermissionDao) Columns() RolePermissionColumns {
	return dao.columns
}

func (dao *RolePermissionDao) Group() string {
	return dao.group
}

func (dao *RolePermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

func (dao *RolePermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
