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
    func main() {
        app := iris.New()
        app.Use(recover.New())
        app.Use(logger.New())
        app.Logger().SetLevel("info")
        router(app)
		uri := fmt.Sprintf("%s:%d",
			config.Conf.Get("app.host").(string),
			config.Conf.Get("app.port").(int64))
        app.Run(iris.Addr(uri), iris.WithoutInterruptHandler)
    }