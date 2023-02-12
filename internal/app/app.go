package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-service/config"
)

func Run(cfg *config.Config) {
	// 初始化gin路由
	gin.SetMode(cfg.RunMode)
	r := gin.Default()

	// 初始化数据库
	//dsn := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ")/" + cfg.DBName + "?parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	r.GET("/welcome", func(c *gin.Context) {
		c.String(http.StatusOK, "config: %s", cfg.DataBase.DBName)
	})
	// 启动项目
	r.Run(":" + cfg.HttpPort)
}
