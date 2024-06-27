package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

type UserColumns struct {
	Id           string //
	Name         string // 用户名
	Avatar       string // 头像
	Password     string //
	UserSalt     string // 加密盐 生成密码用
	Sex          string // 1男 2女
	Status       string // 1正常 2拉黑冻结
	Sign         string // 个性签名
	SecretAnswer string // 密保问题的答案
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
}

var userColumns = UserColumns{
	Id:           "id",
	Name:         "name",
	Avatar:       "avatar",
	Password:     "password",
	UserSalt:     "user_salt",
	Sex:          "sex",
	Status:       "status",
	Sign:         "sign",
	SecretAnswer: "secret_answer",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "user_info",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
