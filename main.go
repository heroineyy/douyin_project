package main

import (
	"byte_douyin_project/config"
	"byte_douyin_project/router"
	"fmt"
	"github.com/swaggo/files" // 注意这里的导入路径
	"github.com/swaggo/gin-swagger"
)

// @title douyin api
// @version 1.0
// @description 极简版抖音API
// @termsOfService http://swagger.io/terms/

// @contact.name heroineyy
// @contact.email 93510710@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /douyin
func main() {
	r := router.InitDouyinRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(fmt.Sprintf(":%d", config.Info.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}
