package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

var (
	client *redis.Client
	ctx    = context.Background()
)

// Redis配置参数
var (
	redisHost     = os.Getenv("REDIS_HOST")
	redisPassword = os.Getenv("REDIS_PASSWORD")
)

// newRedisClient 创建Redis客户端连接池
func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0, // 默认数据库

		// 连接池配置
		PoolSize:     30,                // 最大连接数（相当于MaxActive）
		MinIdleConns: 10,                // 最小空闲连接数
		MaxIdleConns: 50,                // 最大空闲连接数（相当于MaxIdle）
		PoolTimeout:  300 * time.Second, // 连接池超时（相当于IdleTimeout）

		// 连接健康检查
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,

		// 自动重试
		MaxRetries:      3,
		MinRetryBackoff: 300 * time.Millisecond,
		MaxRetryBackoff: 500 * time.Millisecond,
	})
}

func init() {
	client = newRedisClient()

	// 验证连接
	if err := client.Ping(ctx).Err(); err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}

// RedisClient : 获取Redis客户端
func RedisClient() *redis.Client {
	return client
}
