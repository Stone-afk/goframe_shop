package order

import (
	"context"
	"goframe_shop/internal/consts"
	"goframe_shop/internal/dao"
	"goframe_shop/internal/model"
	"goframe_shop/internal/service"
	"goframe_shop/utility"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/os/gtime"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

// Add 下单
func (s *sOrder) Add(ctx context.Context, in model.OrderAddInput) (*model.OrderAddOutput, error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	in.Number = utility.GetOrderNum()
	out := &model.OrderAddOutput{}
	//官方建议的事务闭包处理
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//生成主订单
		lastInsertId, err := dao.OrdersRepo.Ctx(ctx).InsertAndGetId(in)
		if err != nil {
			return err
		}
		//生成商品订单
		for _, info := range in.OrderAddGoodsInfos {
			info.OrderId = gconv.Int(lastInsertId)
			_, err = dao.OrderGoodsRepo.Ctx(ctx).Insert(info)
			if err != nil {
				return err
			}
		}
		// 更新商品销量和库存，todo 后期接入消息
		for _, info := range in.OrderAddGoodsInfos {
			//商品增加销量
			_, err = dao.GoodsRepo.Ctx(ctx).
				WherePri(info.GoodsId).
				Increment(dao.GoodsRepo.Columns().Sale, info.Count)
			if err != nil {
				return err
			}
			//商品减少库存
			_, err2 := dao.GoodsRepo.Ctx(ctx).
				WherePri(info.GoodsId).
				Decrement(dao.GoodsRepo.Columns().Stock, info.Count)
			if err2 != nil {
				return err
			}
			//商品规格减少库存
			_, err3 := dao.GoodsOptionsRepo.Ctx(ctx).WherePri(info.GoodsOptionsId).Decrement(dao.GoodsOptionsRepo.Columns().Stock, info.Count)
			if err3 != nil {
				return err
			}
		}
		out.Id = uint(lastInsertId)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *sOrder) Get(ctx context.Context, in model.OrderDetailInput) (out *model.OrderDetailOutput, err error) {
	err = dao.OrdersRepo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	return
}

func (s *sOrder) List(ctx context.Context, in model.OrderListInput) (out *model.OrderListOutput, err error) {
	//1.获得*gdb.Model对象，方面后续调用
	whereCondition := s.orderListCondition(in)
	m := dao.OrdersRepo.Ctx(ctx).Where(whereCondition)
	//2. 实例化响应结构体
	out = &model.OrderListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//3. 分页查询
	listModel := m.Page(in.Page, in.Size)
	//4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	//5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.OrderListOutputItem, 0, in.Size)
	//6. 把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// todo 优化这里的代码
func (s *sOrder) orderListCondition(in model.OrderListInput) *gmap.Map {
	m := gmap.New()

	if in.Number != "" {
		m.Set(dao.OrdersRepo.Columns().Number+" like ", "%"+in.Number+"%")
	}

	if in.UserId != 0 {
		m.Set(dao.OrdersRepo.Columns().UserId, in.UserId)
	}

	if in.PayType != 0 {
		m.Set(dao.OrdersRepo.Columns().PayType, in.PayType)
	}

	if in.PayAtGte != "" {
		m.Set(dao.OrdersRepo.Columns().PayAt+" >= ", gtime.New(in.PayAtGte).StartOfDay())
	}

	if in.PayAtLte != "" {
		m.Set(dao.OrdersRepo.Columns().PayAt+" <= ", gtime.New(in.PayAtLte).EndOfDay())
	}

	if in.Status != 0 {
		m.Set(dao.OrdersRepo.Columns().Status, in.Status)
	}

	if in.ConsigneePhone != "" {
		m.Set(dao.OrdersRepo.Columns().ConsigneePhone+" like ", "%"+in.ConsigneePhone+"%")
	}

	if in.PriceGte != 0 {
		m.Set(dao.OrdersRepo.Columns().Price+" >= ", in.PriceGte)
	}

	if in.PriceLte != 0 {
		m.Set(dao.OrdersRepo.Columns().Price+" <= ", in.PriceLte)
	}

	if in.DateGte != "" {
		m.Set(dao.OrdersRepo.Columns().CreatedAt+" >= ", gtime.New(in.DateGte).StartOfDay())
	}

	if in.DateLte != "" {
		m.Set(dao.OrdersRepo.Columns().CreatedAt+" <= ", gtime.New(in.DateLte).EndOfDay())
	}

	return m
}
