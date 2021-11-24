package news

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerNews(c *gin.Context) {
	// TODO: 返回爬虫数据的逻辑
	c.JSON(http.StatusOK, gin.H{"name": "index"})
}
