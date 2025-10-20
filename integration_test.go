package go_cache

import (
	"testing"
	"time"
)

func TestAllCacheImplementations(t *testing.T) {
	// 测试所有缓存实现
	caches := map[string]Cache{
		"redis":  NewRedisCache("localhost:6379", "", 0),
		"memory": NewMemoryCache(),
	}

	// 添加文件缓存（需要特殊处理）
	fileCache, err := NewFileCache("./test_integration_cache")
	if err != nil {
		t.Fatalf("创建文件缓存失败: %v", err)
	}
	caches["file"] = fileCache

	// 添加组合缓存
	memoryCache := NewMemoryCache()
	multiCache := NewMultiCache(memoryCache, fileCache)
	caches["multi"] = multiCache

	// 测试所有缓存实现
	for name, cache := range caches {
		t.Run(name, func(t *testing.T) {
			key := "test_key"
			value := "test_value"
			expiration := 5 * time.Second

			// 测试设置键值对
			err := cache.Set(key, value, expiration)
			if err != nil {
				t.Fatalf("设置键值对失败: %v", err)
			}

			// 测试获取键值对
			got, err := cache.Get(key)
			if err != nil {
				t.Fatalf("获取键值对失败: %v", err)
			}

			if got != value {
				t.Errorf("期望值 %s, 实际值 %s", value, got)
			}

			// 测试检查键是否存在
			exists, err := cache.Exists(key)
			if err != nil {
				t.Fatalf("检查键存在失败: %v", err)
			}
			if !exists {
				t.Error("期望键存在，但检查结果为不存在")
			}

			// 测试TTL
			ttl, err := cache.TTL(key)
			if err != nil {
				t.Fatalf("获取TTL失败: %v", err)
			}
			if ttl <= 0 && ttl != -1 {
				t.Error("期望TTL大于0或等于-1（永不过期）")
			}

			// 测试删除键
			err = cache.Delete(key)
			if err != nil {
				t.Fatalf("删除键失败: %v", err)
			}

			// 再次检查键是否存在
			exists, err = cache.Exists(key)
			if err != nil {
				t.Fatalf("检查键存在失败: %v", err)
			}
			if exists {
				t.Error("期望键不存在，但检查结果为存在")
			}

			// 关闭缓存
			cache.Close()
		})
	}
}