package cmd

import (
	"context"
	"goframe_shop/internal/consts"
	"goframe_shop/internal/controller/admin"
	"goframe_shop/internal/controller/goods"
	"goframe_shop/internal/controller/login"
	"goframe_shop/internal/controller/rotation"
	"goframe_shop/internal/controller/user"
	"goframe_shop/internal/service"

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
			// 启动管理后台gtoken
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}
			//管理后台路由组
			s.Group("/api/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//不需要登录的路由组绑定
				group.Bind(
					admin.NewV1().Add, // 管理员
					login.NewV1(),     // 登录
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						hello.NewV1(),
						goods.NewV1(),
						user.NewV1(),
						rotation.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
