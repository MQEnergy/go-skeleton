package vars

import (
	"path"
	"runtime"

	"github.com/MQEnergy/go-skeleton/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	BasePath string              // 根目录
	DB       *gorm.DB            // Mysql数据库
	MDB      map[string]*gorm.DB // mysql多数据库操作
	Redis    *redis.Client       // redis连接池
	Router   *fiber.Router       // 路由
	Config   *config.Config      // 配置
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(path.Dir(filename)))
	BasePath = root
}
