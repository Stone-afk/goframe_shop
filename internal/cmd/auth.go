package cmd

import (
	"context"
	v1 "goframe_shop/api/login/v1"
	"goframe_shop/internal/consts"
	"goframe_shop/internal/dao"
	"goframe_shop/internal/model/entity"
	"goframe_shop/utility"
	"goframe_shop/utility/response"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
)

func loginFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()
	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}

	//验证账号密码是否正确
	adminEntity := entity.Admin{}
	err := dao.AdminRepo.Ctx(ctx).Where(dao.AdminRepo.Columns().Name, name).Scan(&adminEntity)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, adminEntity.UserSalt) != adminEntity.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}

	// 唯一标识，扩展参数user data
	return consts.GTokenAdminPrefix + strconv.Itoa(adminEntity.Id), adminEntity
}

// for 前台项目
func loginFuncFrontend(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}

	//验证账号密码是否正确
	user := entity.User{}
	err := dao.UserRepo.Ctx(ctx).Where(dao.UserRepo.Columns().Name, name).Scan(&user)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, user.UserSalt) != user.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return consts.GTokenFrontendPrefix + strconv.Itoa(user.Id), user
}

// 自定义的登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
		//根据id获得登录用户其他信息
		adminEntity := entity.Admin{}
		err := dao.AdminRepo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminEntity)
		if err != nil {
			return
		}
		//通过角色查询权限
		//先通过角色查询权限id
		var rolePermissions []entity.RolePermission
		err = dao.RolePermissionRepo.Ctx(context.TODO()).
			WhereIn(dao.RolePermissionRepo.Columns().RoleId, g.Slice{adminEntity.RoleIds}).
			Scan(&rolePermissions)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissions {
			permissionIds = append(permissionIds, info.PermissionId)
		}
		var permissions []entity.Permission
		err = dao.PermissionRepo.Ctx(context.TODO()).WhereIn(dao.PermissionRepo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &v1.LoginRes{
			Type:        consts.TokenType,
			Token:       respData.GetString("token"),
			ExpireIn:    consts.GTokenExpireIn, //单位秒,
			IsAdmin:     adminEntity.IsAdmin,
			RoleIds:     adminEntity.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}
