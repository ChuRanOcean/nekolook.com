package main

import (
	"fmt"
	"net/http"

	"nekolook/backend/internal/api"
	"nekolook/backend/pkg/cache"
	"nekolook/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config := LoadConfig()

	// 初始化数据库
	db, err := database.NewPostgreSQL(config.DB)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %v", err))
	}

	// 初始化Redis
	redisClient := cache.NewRedisClient(config.Redis)

	// 创建Gin引擎
	r := gin.Default()

	// 注册中间件
	r.Use(
		database.InjectDB(db),
		cache.InjectRedis(redisClient),
	)

	// 注册路由
	api.RegisterRoutes(r)

	// 启动服务器
	fmt.Printf("NekoLook API Server is running on :%d\n", config.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), r)
}

// LoadConfig 加载配置文件
func LoadConfig() *config.AppConfig {
	// TODO: 实现配置加载
	return &config.AppConfig{
		DB: config.DBConfig{
			Host:     "postgres",
			Port:     5432,
			User:     "neko_admin",
			Password: "meowmeow123",
			DBName:   "nekolook",
		},
		Redis: config.RedisConfig{
			Host: "redis",
			Port: 6379,
		},
		Server: config.ServerConfig{
			Port: 8080,
		},
	}
}
