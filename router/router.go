package router

import (
	"github.com/kataras/iris/v12"
	"lookarm/api"
	"lookarm/config"
	"lookarm/middleware"
)

func InitRouter() {
	app := iris.Default()
	app.UseRouter(middleware.Cors())
	//app.HandleDir("/","web/lookarm/dist")
	//app.HandleDir("/admin","web/admin/dist")
	
	v1 := app.Party("api/v1/")
	v1.Use(middleware.JwtToken())
	{
		// 管理员模块
		v1.Post("adduser", api.AddUser)
		v1.Delete("delete_user/{id:uint}", api.DeleteUser)
		v1.Get("users", api.GetUsers)
		v1.Get("user/info/{id:int}", api.GetUserInfo)
		v1.Put("user/{id:uint}", api.EditUser)
		// 标签管理模块
		v1.Post("tag/add", api.AddTag)
		v1.Put("tag/edit/{id:int}", api.EditTag)
		v1.Delete("tag/delete/{id:int}", api.DeleteTag)
		
		// 分类管理模块
		v1.Post("category/add", api.AddCategory)
		v1.Put("category/edit/{id:int}", api.EditCategory)
		v1.Delete("category/delete/{id:int}", api.DeleteCategory)
		
		// 表单管理模块
		v1.Put("postinfo/edit/{id:int}", api.EditPostInfo)
		v1.Delete("postinfo/delete/{id:int}", api.DeletePostInfo)
		v1.Get("postinfo/category/{id}", api.GetPostInfoCateList)
		
		// App信息管理模块
		v1.Post("appinfo/add", api.AddAppInfo)
		v1.Put("appinfo/edit/{id:int}", api.EditAppInfo)
		v1.Delete("appinfo/delete/{id:int}", api.DeleteAppInfo)
	}
	
	pub := app.Party("api/v1/")
	{
		// 登录
		pub.Post("login", api.Login)
		
		// 标签
		pub.Get("tag/list", api.GetTagList)
		pub.Get("tag/info/{id:int}", api.GetTagInfo)
		
		// 分类
		pub.Get("category/list", api.GetCategoryList)
		pub.Get("category/info/{id:int}", api.GetCategory)
		
		// 表单
		pub.Post("postinfo/add", api.PostAppInfo)
		pub.Get("postinfo/list", api.GetPostAppInfoList)
		pub.Get("postinfo/info/{id:int}", api.GetPostInfo)
		
		// app信息
		pub.Get("appinfo/list", api.GetAppInfoList)
		pub.Get("appinfo/info",api.SearchAppInfo)
		pub.Get("appinfo/info/{id:int}", api.GetAppInfo)
		pub.Get("appinfo/category/{id:int}", api.GetAppInfoCateList)
	}
	_ = app.Listen(config.ServerPort)
	
}
