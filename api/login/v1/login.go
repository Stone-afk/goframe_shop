package v1

import (
	"goframe_shop/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginDoReq struct {
	//g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

// LoginDoRes for jwt
type LoginDoRes struct {
	//Info interface{} `json:"info"`
	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// LoginRes for gtoken
type LoginRes struct {
	Type        string              `json:"type"`
	Token       string              `json:"token"`
	ExpireIn    int                 `json:"expire_in"`
	IsAdmin     int                 `json:"is_admin"`    //是否超管
	RoleIds     string              `json:"role_ids"`    //角色
	Permissions []entity.Permission `json:"permissions"` //权限列表
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type LogoutRes struct {
}
