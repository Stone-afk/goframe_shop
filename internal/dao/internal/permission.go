package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionDao is the data access object for table permission_info.
type PermissionDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns PermissionColumns // columns contains all the column names of Table for convenient usage.
}

type PermissionColumns struct {
	Id        string //
	Name      string // 权限名称
	Path      string // 路径
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// PermissionColumns holds the columns for table permission_info.
var permissionColumns = PermissionColumns{
	Id:        "id",
	Name:      "name",
	Path:      "path",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPermissionDao creates and returns a new DAO object for table data access.
func NewPermissionDao() *PermissionDao {
	return &PermissionDao{
		group:   "default",
		table:   "permission_info",
		columns: permissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PermissionDao) Columns() PermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

func (dao *PermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
