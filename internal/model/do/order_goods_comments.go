// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderGoodsComments is the golang structure of table order_goods_comments for DAO operations like Where/Data.
type OrderGoodsComments struct {
	g.Meta         `orm:"table:order_goods_comments, do:true"`
	Id             interface{} //
	OrderId        interface{} // 订单id
	GoodsId        interface{} // 商品id
	GoodsOptionsId interface{} // 商品规格id
	ParentId       interface{} // 父级评论id
	Content        interface{} // 评论内容
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
}
