package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerIndex(c *gin.Context) {
	// TODO: 返回爬虫数据的逻辑
	c.IndentedJSON(http.StatusOK, gin.H{"name": "index"})
}
