package controller

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"myapp/cmsProject/model"
	"myapp/cmsProject/service"
	"myapp/cmsProject/utils"
)

type AdminController struct {
	//iris框架自动为每个请求都绑定上下文对象
	Ctx iris.Context
	//admin功能实体
	Service service.AdminService
	//session对象
	Session *sessions.Session
}

const (
	ADMINTABLENAME = "admin"
	ADMIN          = "admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//管理员登录功能
//接口：/admin/login
//请求：post
func (ac *AdminController) PostLogin(context iris.Context) mvc.Result{
	var adminLogin AdminLogin
	ac.Ctx.ReadJSON(&adminLogin)

	//数据参数检验
	if adminLogin.UserName==""||adminLogin.Password==""{
		return mvc.Response{
			Object: map[string]interface{}{
				"status":"0",
				"success":"登录失败",
				"message":"用户名或密码为空，请重新填写后尝试登陆",
			},
		}
	}

	//根据用户名、密码到数据库中查询相应的管理信息
	admin, exit:=ac.Service.GetByAdminNameAndPassword(adminLogin.UserName,adminLogin.Password)

	if !exit{
		return mvc.Response{
			Object: map[string]interface{}{
				"status":"0",
				"success":"登录失败",
				"message":"用户名或密码错误，请重新登陆",
			},
		}
	}

	//管理员存在，设置session
	userByte,_:=json.Marshal(admin)
	ac.Session.Set(ADMIN,userByte)

	return mvc.Response{
		Object: map[string]interface{}{
			"status":"1",
			"success":"登录成功",
			"message":"管理员登陆成功",
		},
	}
}

//管理员信息查询接口
//接口：/admin/info
//请求：get
func (ac *AdminController) GetInfo() mvc.Result{
	//从session中获取信息
	userbyte:=ac.Session.Get(ADMIN)

	//若session为空
	if userbyte==nil{
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_UNLOGIN,
				"type": utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	//若session非空，解析数据到admin结构体
	var admin model.Admin
	err:=json.Unmarshal(userbyte.([]byte),&admin)
	if err!=nil{
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_UNLOGIN,
				"type": utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"data": admin.AdminToRespDesc(),
		},
	}

}