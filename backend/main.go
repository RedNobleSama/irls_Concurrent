package main

import "github.com/kataras/iris"

func main() {
	// 1. 创建iris 实例
	app := iris.New()

	// 2. 设置错误模式, 在mvc模式下提示错误
	app.Logger().SetLevel("debug")

	// 3. 注册模板
	template := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)

	// 4. 设置模板目标
	app.HandleDir("assets", iris.Dir("./backend/web/assets"))
	
	// 5. 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(context iris.Context) {
		context.ViewData("message", context.Values().GetStringDefault("message","访问的页面出错"))
		context.ViewLayout("")
		context.View("shared/error.html")
	})

	// 6.注册控制器

	// 7. 启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithOptimizations,
	)
}
