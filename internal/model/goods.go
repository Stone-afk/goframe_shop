package model

import (
	"goframe_shop/internal/model/do"
)

// GoodsUpdateInput 修改内容
type GoodsUpdateInput struct {
	GoodsCreateUpdateBase
	Id uint `json:"id"`
}

type GoodsDeleteInput struct {
	Id uint `json:"id"`
}

type GoodsDeleteOutput struct{}

type GoodsCreateInput struct {
	GoodsCreateUpdateBase
}

type GoodsCreateUpdateBase struct {
	PicUrl           string
	Name             string
	Price            int
	Level1CategoryId int
	Level2CategoryId int
	Level3CategoryId int
	Brand            string
	Stock            int
	Sale             int
	Tags             string
	DetailInfo       string
}

type GoodsCreateOutput struct {
	Id uint `json:"id"`
}

type GoodsGetInput struct {
	Id uint `json:"id" v:"required"`
}

type GoodsGetOutput struct {
	do.Goods
}
