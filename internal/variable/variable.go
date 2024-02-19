package variable

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go-skeleton/pkg/config"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

var (
	BasePath string         // 根目录
	DB       *gorm.DB       // Mysql数据库
	Log      *slog.Logger   // 日志
	Redis    *redis.Client  // redis连接池
	Router   *fiber.Router  // 路由
	Config   *config.Config // 配置
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	BasePath = dir
}
