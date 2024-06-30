// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsOptions is the golang structure for table goods_options.
type GoodsOptions struct {
	Id        int         `json:"id"        orm:"id"         description:"ID"`
	GoodsId   int         `json:"goodsId"   orm:"goods_id"   description:"商品id"`
	PicUrl    string      `json:"picUrl"    orm:"pic_url"    description:"图片"`
	Name      string      `json:"name"      orm:"name"       description:"商品名称"`
	Price     int         `json:"price"     orm:"price"      description:"价格 单位分"`
	Stock     int         `json:"stock"     orm:"stock"      description:"库存"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
