package src

import "github.com/gin-gonic/gin"

func GetTopicDetail(ctx *gin.Context) {
	ctx.String(200, "topic_id", ctx.Param("topic_id"))
}