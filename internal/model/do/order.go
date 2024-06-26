package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Order struct {
	g.Meta           `orm:"table:order_info, do:true"`
	Id               interface{} //
	Number           interface{} // 订单编号
	UserId           interface{} // 用户id
	PayType          interface{} // 支付方式 1微信 2支付宝 3云闪付
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	Remark           interface{} // 备注
	PayAt            *gtime.Time // 支付时间
	Status           interface{} // 订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价
	ConsigneeName    interface{} // 收货人姓名
	ConsigneePhone   interface{} // 收货人手机号
	ConsigneeAddress interface{} // 收货人详细地址
	Price            interface{} // 订单金额 单位分
	CouponPrice      interface{} // 优惠券金额 单位分
	ActualPrice      interface{} // 实际支付金额 单位分
}
