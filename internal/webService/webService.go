package webService

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaogogonuo/cct-spider/internal/index"
)

func RunService() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.POST("/spider/post", func(c *gin.Context) {
		var f []index.Field
		err := c.ShouldBind(&f)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(f)
	})
	if err := r.Run("0.0.0.0:8888"); err != nil {
		panic(err)
	}
}
