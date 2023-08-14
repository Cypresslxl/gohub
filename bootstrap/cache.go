// Package bootstrap 启动程序功能
package bootstrap

import (
	"fmt"
	"gohub/pkg/cache"
	"gohub/pkg/config"
)

// SetUpCache 缓存
func SetupCache() {
	//	初始化缓存专用的 redis client，使用专属缓存 DB
	redis := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)

	cache.InitWithCacheStore(redis)

}
