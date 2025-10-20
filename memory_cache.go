package main

import (
	"sync"
	"time"
)

// MemoryCache 实现了基于内存的缓存
type MemoryCache struct {
	data map[string]*cacheItem
	mu   sync.RWMutex
	stop chan bool
}

// cacheItem 表示缓存中的一个项目
type cacheItem struct {
	value      string
	expiration time.Time
}

// NewMemoryCache 创建一个新的内存缓存实例
func NewMemoryCache() *MemoryCache {
	cache := &MemoryCache{
		data: make(map[string]*cacheItem),
		stop: make(chan bool),
	}

	// 启动过期清理协程
	go cache.cleanup()

	return cache
}

// Set 将键值对存储到缓存中，并设置过期时间
func (m *MemoryCache) Set(key string, value interface{}, expiration time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var expirationTime time.Time
	if expiration > 0 {
		expirationTime = time.Now().Add(expiration)
	}

	m.data[key] = &cacheItem{
		value:      value.(string),
		expiration: expirationTime,
	}

	return nil
}

// Get 从缓存中获取指定键的值
func (m *MemoryCache) Get(key string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	item, exists := m.data[key]
	if !exists {
		return "", ErrKeyNotFound
	}

	// 检查是否过期
	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		return "", ErrKeyNotFound
	}

	return item.value, nil
}

// Delete 从缓存中删除指定键
func (m *MemoryCache) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.data, key)
	return nil
}

// Exists 检查指定键是否存在于缓存中
func (m *MemoryCache) Exists(key string) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	item, exists := m.data[key]
	if !exists {
		return false, nil
	}

	// 检查是否过期
	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		return false, nil
	}

	return true, nil
}

// Expire 设置键的过期时间
func (m *MemoryCache) Expire(key string, expiration time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	item, exists := m.data[key]
	if !exists {
		return ErrKeyNotFound
	}

	if expiration > 0 {
		item.expiration = time.Now().Add(expiration)
	} else {
		item.expiration = time.Time{}
	}

	return nil
}

// TTL 获取键的剩余生存时间
func (m *MemoryCache) TTL(key string) (time.Duration, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	item, exists := m.data[key]
	if !exists {
		return 0, ErrKeyNotFound
	}

	// 检查是否过期
	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		return 0, ErrKeyNotFound
	}

	if item.expiration.IsZero() {
		// 永不过期
		return -1, nil
	}

	return time.Until(item.expiration), nil
}

// Close 关闭缓存连接
func (m *MemoryCache) Close() error {
	close(m.stop)
	return nil
}

// cleanup 定期清理过期的缓存项
func (m *MemoryCache) cleanup() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			m.mu.Lock()
			now := time.Now()
			for key, item := range m.data {
				if !item.expiration.IsZero() && now.After(item.expiration) {
					delete(m.data, key)
				}
			}
			m.mu.Unlock()
		case <-m.stop:
			return
		}
	}
}
