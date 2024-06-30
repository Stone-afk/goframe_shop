// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderGoods is the golang structure for table order_goods.
type OrderGoods struct {
	Id             int         `json:"id"             orm:"id"               description:"商品维度的订单表"`
	OrderId        int         `json:"orderId"        orm:"order_id"         description:"关联的主订单表"`
	GoodsId        int         `json:"goodsId"        orm:"goods_id"         description:"商品id"`
	GoodsOptionsId int         `json:"goodsOptionsId" orm:"goods_options_id" description:"商品规格id"`
	Count          int         `json:"count"          orm:"count"            description:"商品数量"`
	PayType        int         `json:"payType"        orm:"pay_type"         description:"支付方式 1微信 2支付宝 3云闪付"`
	Remark         string      `json:"remark"         orm:"remark"           description:"备注"`
	Status         int         `json:"status"         orm:"status"           description:"订单状态 0待支付 1已支付 3已确认收货"`
	Price          int         `json:"price"          orm:"price"            description:"订单金额 单位分"`
	CouponPrice    int         `json:"couponPrice"    orm:"coupon_price"     description:"优惠券金额 单位分"`
	ActualPrice    int         `json:"actualPrice"    orm:"actual_price"     description:"实际支付金额 单位分"`
	PayAt          *gtime.Time `json:"payAt"          orm:"pay_at"           description:"支付时间"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`
}
