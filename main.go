package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"myapp/cmsProject/config"
	"myapp/cmsProject/controller"
	"myapp/cmsProject/datasource"
	"myapp/cmsProject/service"
	"time"
)

func main(){
	app:=newApp()

	//应用App配置
	configation(app)

	//路由配置
	mvcHandle(app)

	config:=config.InitConfig()
	addr:=":"+config.Port

	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}

//构建App
func newApp() *iris.Application{
	app:=iris.New()

	//设置日志级别，开发阶段为debug
	app.Logger().SetLevel("debug")

	//注册静态资源
	app.HandleDir("/static","./static")
	app.HandleDir("/manage/static","./static")

	//注册视图文件
	app.RegisterView(iris.HTML("./static",".html"))
	//初始界面
	app.Get("/", func(context iris.Context) {
		context.View("index.html")
	})
	return app
}

//项目配置
func configation(app *iris.Application){
	//配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	//未发现的错误
	app.OnErrorCode(iris.StatusNotFound, func(context iris.Context) {
		context.JSON(iris.Map{
			"errmsg":iris.StatusNotFound,
			"msg":"not found",
			"data": iris.Map{},
		})
	})

	//内部错误
	app.OnErrorCode(iris.StatusInternalServerError, func(context iris.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg": "internal error",
			"data": iris.Map{},
		})
	})
}

//路由配置 MVC架构模式处理
func mvcHandle(app *iris.Application){
	//启用session
	sessManager:=sessions.New(sessions.Config{
		Cookie: "sessioncookie",
		Expires: 24*time.Hour,
	})

	engine:=datasource.NewMysqlEngine()

	//管理员模块功能
	adminService:=service.NewAdminService(engine)

	admin:=mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessManager.Start,
		)
	admin.Handle(new(controller.AdminController))
}
