package hello

import (
	"context"
	v1 "goframe_shop/api/backend/hello/v1"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
