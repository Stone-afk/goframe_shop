package v1

import "github.com/gogf/gf/v2/frame/g"

type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete" tags:"商品" method:"delete" summary:"删除商品"`
	Id     uint `v:"min:1#请选择需要删除的商品" dc:"商品id"`
}

type GoodsDeleteRes struct{}

type GoodsAddReq struct {
	g.Meta `path:"/goods/add" tags:"商品" method:"post" summary:"添加商品"`
	GoodsCommonAddUpdate
}

type GoodsCommonAddUpdate struct {
	PicUrl           string `json:"pic_url"           description:"图片"`
	Name             string `json:"name"             description:"商品名称" v:"required#名称必传"`
	Price            int    `json:"price"            description:"价格 单位分" v:"required#价格必传"`
	Level1CategoryId uint   `json:"level1_category_id" description:"1级分类id"`
	Level2CategoryId uint   `json:"level2_category_id" description:"2级分类id"`
	Level3CategoryId uint   `json:"level3_category_id" description:"3级分类id"`
	Brand            string `json:"brand"            description:"品牌" v:"max-length:30#品牌名称最大30个字"`
	Stock            int    `json:"stock"            description:"库存"`
	Sale             uint   `json:"sale"             description:"销量"`
	Tags             string `json:"tags"             description:"标签"`
	DetailInfo       string `json:"detail_info"       description:"商品详情"`
}

type GoodsAddRes struct {
	Id uint `json:"id"`
}

type GoodsGetReq struct {
	g.Meta `path:"/goods/detail" method:"get" tags:"商品" summary:"商品详情接口"`
	Id     uint `json:"id" v:"required"`
}

type GoodsGetRes struct {
	*GoodsCommonAddUpdate
}