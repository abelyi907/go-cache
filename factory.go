package go_cache

// CacheType 定义缓存类型
type CacheType string

const (
	// RedisCacheType Redis缓存类型
	RedisCacheType CacheType = "redis"

	// MemoryCacheType 内存缓存类型
	MemoryCacheType CacheType = "memory"

	// FileCacheType 文件缓存类型
	FileCacheType CacheType = "file"
)

// CacheConfig 缓存配置
type CacheConfig struct {
	// Type 缓存类型
	Type CacheType

	// Redis配置
	RedisAddr     string
	RedisPassword string
	RedisDB       int

	// File配置
	FileDir string

	PrefixKey string // 缓存key的前缀
}

// NewCache 根据配置创建缓存实例
func NewCache(config CacheConfig) (Cache, error) {
	switch config.Type {
	case RedisCacheType:
		return NewRedisCache(config.RedisAddr, config.RedisPassword, config.RedisDB, config.PrefixKey), nil
	case MemoryCacheType:
		return NewMemoryCache(), nil
	case FileCacheType:
		return NewFileCache(config.FileDir)
	default:
		return NewMemoryCache(), nil
	}
}
