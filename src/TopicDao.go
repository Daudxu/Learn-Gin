package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MustLogin() gin.HandlerFunc{
	return func (ctx *gin.Context)  {
		if _, status := ctx.GetQuery("token");!status {
            ctx.String(http.StatusUnauthorized, "缺少token参数")
			ctx.Abort()
		}else{
			ctx.Next()
		}
	}
}

func GetTopicDetail(ctx *gin.Context) {
	// ctx.String(200, "topic_id", ctx.Param("topic_id"))
	ctx.JSON(200, CreateTopic(1,"ceasdasdasd"))
}

func NewTopic(ctx *gin.Context) {
	topic := Topic{} 
	err := ctx.BindJSON(&topic)
	if err != nil {
		ctx.String(400, "参数错误：%s", err.Error())
	}else{
		ctx.JSON(200, topic)
	}
}

func DelTopic(ctx *gin.Context) {
	ctx.String(200, "topic_id", "删除帖子")
}

func GetTopList(ctx *gin.Context) {
    // quick
   query := TopicQuery{}
   err := ctx.BindQuery(&query)
   if(err != nil) {
	ctx.String(400, "参数错误：%s", err.Error())
   }else{
	ctx.JSON(200, query)
   }
}
