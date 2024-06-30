// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsOptions is the golang structure of table goods_options for DAO operations like Where/Data.
type GoodsOptions struct {
	g.Meta    `orm:"table:goods_options, do:true"`
	Id        interface{} // ID
	GoodsId   interface{} // 商品id
	PicUrl    interface{} // 图片
	Name      interface{} // 商品名称
	Price     interface{} // 价格 单位分
	Stock     interface{} // 库存
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
