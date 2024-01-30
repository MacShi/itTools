 package main

 import (
	 "fmt"
	 "github.com/kataras/iris"
	 "github.com/kataras/iris/middleware/logger"
	 "github.com/kataras/iris/middleware/recover"
	 "itTools/config"
	 "itTools/control"
 )

 func router(app  *iris.Application)  {
      //app.Get("/", func(ctx iris.Context) {
      //    ctx.HTML(" <h1>hi, I just exist in order to see if the server is closed</h1>")
      //})
	 app.Get("/", func(ctx iris.Context) {
	 	ctx.HTML("<h1>welcole</h1>")

	 })
      apis:=app.Party("/apis")
      {
      	apis.Post("/hide",control.HideStr)
      	apis.Get("/clientInfo",control.ClientInfo)
      	apis.Post("/getPwd",control.GetPwd)
      	apis.Post("/upload",control.UploadFile)
	 }

 }
 //如果使用身份验证，可以在这统一进行
 //func before(ctx iris.Context) {
 	//通过检查，接着继续操作，不通过直接拒绝
	// aa :=ctx.GetHeader("aaa")
	// if aa=="11"{
	// 	ctx.Next()
	// }else {
	// 	ctx.JSON(iris.StatusBadGateway)
	// }
 //}
    func main() {
        app := iris.New()
        app.Use(recover.New())
        app.Use(logger.New())
        app.Logger().SetLevel("info")
        //全局注册before
        //app.UseGlobal(before)
        router(app)
		uri := fmt.Sprintf("%s:%d",
			config.Conf.Get("app.host").(string),
			config.Conf.Get("app.port").(int64))
        app.Run(iris.Addr(uri), iris.WithoutInterruptHandler)
    }