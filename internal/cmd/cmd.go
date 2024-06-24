package cmd

import (
	"context"
	"goframe_shop/internal/consts"
	"goframe_shop/internal/controller/goods"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"goframe_shop/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			//管理后台路由组
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(ghttp.MiddlewareHandlerResponse)
					group.Bind(
						hello.NewV1(),
						goods.NewV1(), //商品管理
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
