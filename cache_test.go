package go_cache

import (
	"testing"
	"time"
)

func TestRedisCache_SetAndGet(t *testing.T) {
	cache := redisServer

	key := "test_key"
	value := "test_value"
	expiration := 500 * time.Second

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
}

func TestRedisCache_Delete(t *testing.T) {
	cache := redisServer

	key := "test_delete_key"
	value := "test_value"
	expiration := 5 * time.Second

	// 设置键值对
	err := cache.Set(key, value, expiration)
	if err != nil {
		t.Fatalf("设置键值对失败: %v", err)
	}

	// 删除键
	err = cache.Delete(key)
	if err != nil {
		t.Fatalf("删除键失败: %v", err)
	}

	// 尝试获取已删除的键
	_, err = cache.Get(key)
	if err == nil {
		t.Error("期望键已被删除，但获取成功")
	}
}

func TestRedisCache_Exists(t *testing.T) {
	cache := redisServer

	key := "test_exists_key"
	value := "test_value"
	expiration := 5 * time.Second

	// 检查不存在的键
	exists, err := cache.Exists(key)
	if err != nil {
		t.Fatalf("检查键存在失败: %v", err)
	}
	if exists {
		t.Error("期望键不存在，但检查结果为存在")
	}

	// 设置键值对
	err = cache.Set(key, value, expiration)
	if err != nil {
		t.Fatalf("设置键值对失败: %v", err)
	}

	// 检查存在的键
	exists, err = cache.Exists(key)
	if err != nil {
		t.Fatalf("检查键存在失败: %v", err)
	}
	if !exists {
		t.Error("期望键存在，但检查结果为不存在")
	}
}

func TestRedisCache_Expire(t *testing.T) {
	cache := redisServer

	key := "test_expire_key"
	value := "test_value"
	expiration := 1 * time.Second

	// 设置键值对
	err := cache.Set(key, value, expiration)
	if err != nil {
		t.Fatalf("设置键值对失败: %v", err)
	}

	// 设置新的过期时间
	newExpiration := 10 * time.Second
	err = cache.Expire(key, newExpiration)
	if err != nil {
		t.Fatalf("设置过期时间失败: %v", err)
	}

	// 检查TTL
	ttl, err := cache.TTL(key)
	if err != nil {
		t.Fatalf("获取TTL失败: %v", err)
	}

	if ttl <= 0 {
		t.Error("期望TTL大于0")
	}
}

func TestMemoryCache_SetAndGet(t *testing.T) {
	cache := NewMemoryCache()
	defer cache.Close()

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
}

func TestMemoryCache_Delete(t *testing.T) {
	cache := NewMemoryCache()
	defer cache.Close()

	key := "test_delete_key"
	value := "test_value"
	expiration := 5 * time.Second

	// 设置键值对
	err := cache.Set(key, value, expiration)
	if err != nil {
		t.Fatalf("设置键值对失败: %v", err)
	}

	// 删除键
	err = cache.Delete(key)
	if err != nil {
		t.Fatalf("删除键失败: %v", err)
	}

	// 尝试获取已删除的键
	_, err = cache.Get(key)
	if err == nil {
		t.Error("期望键已被删除，但获取成功")
	}
}

func TestMemoryCache_Exists(t *testing.T) {
	cache := NewMemoryCache()
	defer cache.Close()

	key := "test_exists_key"
	value := "test_value"
	expiration := 5 * time.Second

	// 检查不存在的键
	exists, err := cache.Exists(key)
	if err != nil {
		t.Fatalf("检查键存在失败: %v", err)
	}
	if exists {
		t.Error("期望键不存在，但检查结果为存在")
	}

	// 设置键值对
	err = cache.Set(key, value, expiration)
	if err != nil {
		t.Fatalf("设置键值对失败: %v", err)
	}

	// 检查存在的键
	exists, err = cache.Exists(key)
	if err != nil {
		t.Fatalf("检查键存在失败: %v", err)
	}
	if !exists {
		t.Error("期望键存在，但检查结果为不存在")
	}
}

func TestFileCache_SetAndGet(t *testing.T) {
	cache, err := NewFileCache("./test_cache")
	if err != nil {
		t.Fatalf("创建文件缓存失败: %v", err)
	}
	defer cache.Close()

	key := "test_key"
	value := "test_value"
	expiration := 5 * time.Second

	// 测试设置键值对
	err = cache.Set(key, value, expiration)
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
}

func TestFileCache_Delete(t *testing.T) {
	cache, err := NewFileCache("./test_cache")
	if err != nil {
		t.Fatalf("创建文件缓存失败: %v", err)
	}
	defer cache.Close()

	key := "test_delete_key"
	value := "test_value"
	expiration := 5 * time.Second

	// 设置键值对
	err = cache.Set(key, value, expiration)
	if err != nil {
		t.Fatalf("设置键值对失败: %v", err)
	}

	// 删除键
	err = cache.Delete(key)
	if err != nil {
		t.Fatalf("删除键失败: %v", err)
	}

	// 尝试获取已删除的键
	_, err = cache.Get(key)
	if err == nil {
		t.Error("期望键已被删除，但获取成功")
	}
}

func TestMultiCache_SetAndGet(t *testing.T) {
	memoryCache := NewMemoryCache()
	fileCache, _ := NewFileCache("./test_cache")
	cache := NewMultiCache(memoryCache, fileCache)
	defer cache.Close()

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
}

func TestUseOneCacheOfRedis(t *testing.T) {
	cc, err := NewCache(CacheConfig{
		Type:          RedisCacheType,
		RedisAddr:     redisUrl,
		RedisPassword: redisPassword,
		RedisDB:       redisDb,
	})
	if err != nil {
		t.Error(err)
	}
	if cc.Set("test1", "test1", time.Second*120) != nil {
		t.Error("set cache error")
	}
	if cc.Set("test2", "test2", time.Second*120) != nil {
		t.Error("set cache error")
	}
	rsp, err := cc.Get("test2")
	if err != nil {
		t.Error(err)
	}
	if rsp != "test2" {
		t.Error("get cache error")
	}

	if cc.Set("test2", "test2_1", time.Second*120) != nil {
		t.Error("set cache error")
	}
	rsp, err = cc.Get("test2")
	if err != nil {
		t.Error(err)
	}
	if rsp != "test2_1" {
		t.Error("get cache error")
	}

	if cc.Delete("test2") != nil {
		t.Error("delete cache error")
	}

}

func TestUseOneCacheOfFile(t *testing.T) {
	cc, err := NewCache(CacheConfig{
		Type:    FileCacheType,
		FileDir: "./test_cache",
	})
	if err != nil {
		t.Error(err)
	}

	if cc.Set("test1", `{"key1":"我是key1"}`, time.Second*120) != nil {
		t.Error("set cache error")
	}
	if cc.Set("test2", `{"key1":"我是key2"}`, time.Second*120) != nil {
		t.Error("set cache error")
	}
	rsp, err := cc.Get("test2")
	if err != nil {
		t.Error(err)
	}
	if rsp != `{"key1":"我是key2"}` {
		t.Error("get cache error")
	}
	if cc.Set("test3", "aa", time.Second*120) != nil {
		t.Error("set cache3 error")
	}

	//
	//if cc.Set("test2", "test2_1", time.Second*120) != nil {
	//	t.Error("set cache error")
	//}
	//rsp, err = cc.Get("test2")
	//if err != nil {
	//	t.Error(err)
	//}
	//if rsp != "test2_1" {
	//	t.Error("get cache error")
	//}

	//if cc.Delete("test2") != nil {
	//	t.Error("delete cache error")
	//}

}
