package go_cache

import (
	"time"
)

// Cache 定义了缓存接口
type Cache interface {
	// Set 将键值对存储到缓存中，并设置过期时间
	Set(key string, value interface{}, expiration time.Duration) error

	// Get 从缓存中获取指定键的值
	Get(key string) (string, error)

	// Delete 从缓存中删除指定键
	Delete(key string) error

	// Exists 检查指定键是否存在于缓存中
	Exists(key string) (bool, error)

	// Expire 设置键的过期时间
	Expire(key string, expiration time.Duration) error

	// TTL 获取键的剩余生存时间
	TTL(key string) (time.Duration, error)

	// Close 关闭缓存连接
	Close() error
}
