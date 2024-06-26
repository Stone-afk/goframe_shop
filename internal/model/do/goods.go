// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure of table goods for DAO operations like Where/Data.
type Goods struct {
	g.Meta           `orm:"table:goods, do:true"`
	Id               interface{} //
	PicUrl           interface{} // 图片
	Name             interface{} // 商品名称
	Price            interface{} // 价格 单位分
	Level1CategoryId interface{} // 1级分类id
	Level2CategoryId interface{} // 2级分类id
	Level3CategoryId interface{} // 3级分类id
	Brand            interface{} // 品牌
	Stock            interface{} // 库存
	Sale             interface{} // 销量
	Tags             interface{} // 标签
	DetailInfo       interface{} // 商品详情
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	DeletedAt        *gtime.Time //
}
