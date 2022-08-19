package main

// go env -w GOOS=linux
// go env -w GOARCH=adm64
import (
	. "gin/src"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1/topics")
	{
		v1.GET("", func(ctx *gin.Context) {
			ctx.String(200,"TEST")
		})

		v1.GET("/:topic_id", GetTopicDetail)
	}
	// logger, err := zap.NewProduction()
	// // file,_:=os.Create("./test.log")
	// // writeSyncer:=zapcore.AddSync(file)
	// if err != nil {
	// 	panic(err)
	// }

	// r.Use(func(ctx *gin.Context) {
	// 	// fmt.Println(ctx.Request)
	//    s := time.Now()
	//    ctx.Next()
	//    logger.Info("incming request",zap.String("path", ctx.Request.URL.Path), zap.String("path", ctx.Writer.Header().Get("requestId")), zap.Int("status", ctx.Writer.Status()), zap.Duration("elapsed", time.Now().Sub(s)))
	// },func(ctx *gin.Context) {
	// 	ctx.Set("requestId",  rand.Int())
	// 	ctx.Next()
	// })

	// r.GET("/ping", func(c *gin.Context) {
	// 	h:= gin.H{
	// 		"message": "pong",
	// 	}
	// 	if rid, exists := c.Get("requestId"); exists {
	// 		h["requestId"] = rid
	// 	}
	// 	c.JSON(200, h)
	// })

	// r.GET("/hello", func(c *gin.Context) {
	// 	c.String(200, "hello")
	// })

	r.Run() 
}