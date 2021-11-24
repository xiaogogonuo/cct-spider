package webService

import (
	"github.com/gin-gonic/gin"

	"github.com/xiaogogonuo/cct-spider/internal/webService/index"
	"github.com/xiaogogonuo/cct-spider/internal/webService/news"
)

func RunService() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	spider := r.Group("/spider")

	{
		spider.GET("/index", index.HandlerIndex)
		spider.GET("/news", news.HandlerNews)
	}

	if err := r.Run("0.0.0.0:8888"); err != nil {
		panic(err)
	}
}
