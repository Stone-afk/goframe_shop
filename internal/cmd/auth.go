package cmd

import (
	"context"
	frontend "goframe_shop/api/frontend/user/v1"
	backend "goframe_shop/api/login/v1"
	"goframe_shop/internal/consts"
	"goframe_shop/internal/dao"
	"goframe_shop/internal/model/entity"
	"goframe_shop/utility"
	"goframe_shop/utility/response"
	"strconv"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
)

// StartBackendGToken 管理后台相关
func StartBackendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	gfAdminToken = &gtoken.GfToken{
		CacheMode:        consts.CacheModeRedis,
		ServerName:       consts.BackendServerName,
		LoginPath:        "/login",
		LoginBeforeFunc:  loginFunc,
		LoginAfterFunc:   loginAfterFunc,
		LogoutPath:       "/user/logout",
		AuthPaths:        g.SliceStr{"/backend/admin/info"},
		AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		AuthAfterFunc:    authAfterFunc,
		MultiLogin:       consts.MultiLogin,
	}
	err = gfAdminToken.Start()
	return
}

// StartFrontendGToken 前台登录gtoken相关
func StartFrontendGToken() (gfFrontendToken *gtoken.GfToken, err error) {
	gfFrontendToken = &gtoken.GfToken{
		CacheMode:       consts.CacheModeRedis,
		ServerName:      consts.BackendServerName,
		LoginPath:       "/login",
		LoginBeforeFunc: loginFuncFrontend,
		LoginAfterFunc:  loginAfterFuncFrontend,
		LogoutPath:      "/user/logout",
		//AuthPaths:        g.SliceStr{"/backend/admin/info"},
		//AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		AuthAfterFunc: authAfterFuncFrontend,
		MultiLogin:    consts.FrontendMultiLogin,
	}
	//todo 去掉全局校验，只用cmd中的路由组校验
	//err = gfAdminToken.Start()
	return
}

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
		data := &backend.LoginRes{
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

// 自定义的登录之后的函数 for前台项目
func loginAfterFuncFrontend(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		userId := gstr.StrEx(userKey, consts.GTokenFrontendPrefix)
		//根据id获得登录用户其他信息
		userInfo := entity.User{}
		err := dao.UserRepo.Ctx(context.TODO()).WherePri(userId).Scan(&userInfo)
		if err != nil {
			return
		}
		data := &frontend.LoginRes{
			Type:     consts.TokenType,
			Token:    respData.GetString("token"),
			ExpireIn: consts.GTokenExpireIn, //单位秒,
		}
		data.Name = userInfo.Name
		data.Avatar = userInfo.Avatar
		data.Sign = userInfo.Sign
		data.Status = uint8(userInfo.Status)
		response.JsonExit(r, 0, "", data)
	}
	return
}

// 登录鉴权中间件for后台
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var adminInfo entity.Admin
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}

// 登录鉴权中间件for前台
func authAfterFuncFrontend(r *ghttp.Request, respData gtoken.Resp) {
	var userInfo entity.User
	err := gconv.Struct(respData.GetString("data"), &userInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	r.SetCtxVar(consts.CtxUserId, userInfo.Id)
	r.SetCtxVar(consts.CtxUserName, userInfo.Name)
	r.SetCtxVar(consts.CtxUserAvatar, userInfo.Avatar)
	r.SetCtxVar(consts.CtxUserSex, userInfo.Sex)
	r.SetCtxVar(consts.CtxUserSign, userInfo.Sign)
	r.SetCtxVar(consts.CtxUserStatus, userInfo.Status)
	r.Middleware.Next()
}
