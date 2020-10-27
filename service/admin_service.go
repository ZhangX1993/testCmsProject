package service

import (
	"github.com/go-xorm/xorm"
	"myapp/cmsProject/model"
)

/*
管理员服务
标准的开发模式将每个实体提供的功能以接口标准的形式定义，供控制层进行调用
 */

type AdminService interface {
	//通过管理员用户名+密码获取管理员实体，如果查询到，返回管理员实体，返回true
	GetByAdminNameAndPassword(username,password string)(model.Admin,bool)
	//获取管理员总数
	GetAdminCount()(int64,error)
}

func NewAdminService(db *xorm.Engine) AdminService{
	return &adminService{
		engine:db,
	}
}

/*
管理员服务实现结构体
 */
type adminService struct {  //adminService实现了AdminService接口
	engine *xorm.Engine
}

//通过用户名和密码查询管理员方法实现
func(ac *adminService) GetByAdminNameAndPassword(username,password string)(model.Admin,bool){
	var admin model.Admin
	ac.engine.Where("user_name=? and pwd=?",username,password).Get(&admin)
	return admin,admin.AdminId!=0
}

//获取管理员总数
func (ac *adminService) GetAdminCount()(int64,error){
	count,err:=ac.engine.Count(new(model.Admin))
	if err!=nil{
		panic(err.Error())
		return 0, err
	}
	return count,nil
}