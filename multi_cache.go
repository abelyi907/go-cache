package go_cache

import (
	"time"
)

// MultiCache 组合多种缓存实现
type MultiCache struct {
	caches []Cache
}

// NewMultiCache 创建一个新的组合缓存实例
func NewMultiCache(caches ...Cache) *MultiCache {
	return &MultiCache{
		caches: caches,
	}
}

// Set 将键值对存储到所有缓存中，并设置过期时间
func (m *MultiCache) Set(key string, value interface{}, expiration time.Duration) error {
	for _, cache := range m.caches {
		err := cache.Set(key, value, expiration)
		if err != nil {
			// 记录错误但继续设置其他缓存
			// 在实际应用中，可能需要更好的错误处理机制
		}
	}
	return nil
}

// Get 从缓存中获取指定键的值，按顺序查找直到找到
func (m *MultiCache) Get(key string) (string, error) {
	for i, cache := range m.caches {
		value, err := cache.Get(key)
		if err == nil {
			// 如果在后面的缓存中找到了，在前面的缓存中设置该值（提升性能）
			for j := 0; j < i; j++ {
				cacheErr := m.caches[j].Set(key, value, 0) // 使用默认过期时间
				if cacheErr != nil {
					// 记录错误但继续
				}
			}
			return value, nil
		}
	}
	return "", ErrKeyNotFound
}

// Delete 从所有缓存中删除指定键
func (m *MultiCache) Delete(key string) error {
	for _, cache := range m.caches {
		err := cache.Delete(key)
		if err != nil {
			// 记录错误但继续删除其他缓存
		}
	}
	return nil
}

// Exists 检查指定键是否存在于任意缓存中
func (m *MultiCache) Exists(key string) (bool, error) {
	for _, cache := range m.caches {
		exists, err := cache.Exists(key)
		if err == nil && exists {
			return true, nil
		}
	}
	return false, nil
}

// Expire 设置所有缓存中键的过期时间
func (m *MultiCache) Expire(key string, expiration time.Duration) error {
	for _, cache := range m.caches {
		err := cache.Expire(key, expiration)
		if err != nil {
			// 记录错误但继续设置其他缓存
		}
	}
	return nil
}

// TTL 获取键的剩余生存时间（从第一个找到的缓存中获取）
func (m *MultiCache) TTL(key string) (time.Duration, error) {
	for _, cache := range m.caches {
		ttl, err := cache.TTL(key)
		if err == nil {
			return ttl, nil
		}
	}
	return 0, ErrKeyNotFound
}

// Close 关闭所有缓存连接
func (m *MultiCache) Close() error {
	for _, cache := range m.caches {
		err := cache.Close()
		if err != nil {
			// 记录错误但继续关闭其他缓存
		}
	}
	return nil
}