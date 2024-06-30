// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderGoodsCommentsDao is the data access object for table order_goods_comments.
type OrderGoodsCommentsDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns OrderGoodsCommentsColumns // columns contains all the column names of Table for convenient usage.
}

// OrderGoodsCommentsColumns defines and stores column names for table order_goods_comments.
type OrderGoodsCommentsColumns struct {
	Id             string //
	OrderId        string // 订单id
	GoodsId        string // 商品id
	GoodsOptionsId string // 商品规格id
	ParentId       string // 父级评论id
	Content        string // 评论内容
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
}

// orderGoodsCommentsColumns holds the columns for table order_goods_comments.
var orderGoodsCommentsColumns = OrderGoodsCommentsColumns{
	Id:             "id",
	OrderId:        "order_id",
	GoodsId:        "goods_id",
	GoodsOptionsId: "goods_options_id",
	ParentId:       "parent_id",
	Content:        "content",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewOrderGoodsCommentsDao creates and returns a new DAO object for table data access.
func NewOrderGoodsCommentsDao() *OrderGoodsCommentsDao {
	return &OrderGoodsCommentsDao{
		group:   "default",
		table:   "order_goods_comments",
		columns: orderGoodsCommentsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderGoodsCommentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderGoodsCommentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderGoodsCommentsDao) Columns() OrderGoodsCommentsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderGoodsCommentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderGoodsCommentsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderGoodsCommentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
