package main

import (
	"fmt"
	"log"
	"time"
)

// ExampleRedis 使用Redis缓存示例
func ExampleRedis() {
	// 创建Redis缓存实例
	// 注意：需要确保Redis服务器正在运行，并且地址正确
	cache := NewRedisCache("localhost:6379", "", 0)
	defer cache.Close()

	// 设置键值对，过期时间10秒
	err := cache.Set("name", "go-cache", 10*time.Second)
	if err != nil {
		log.Fatal("设置缓存失败:", err)
	}

	// 获取键值
	value, err := cache.Get("name")
	if err != nil {
		log.Fatal("获取缓存失败:", err)
	}
	fmt.Println("获取到的值:", value)

	// 检查键是否存在
	exists, err := cache.Exists("name")
	if err != nil {
		log.Fatal("检查键存在失败:", err)
	}
	fmt.Println("键是否存在:", exists)

	// 获取键的剩余生存时间
	ttl, err := cache.TTL("name")
	if err != nil {
		log.Fatal("获取TTL失败:", err)
	}
	fmt.Printf("剩余生存时间: %v\n", ttl)

	// 删除键
	err = cache.Delete("name")
	if err != nil {
		log.Fatal("删除键失败:", err)
	}

	// 再次检查键是否存在
	exists, err = cache.Exists("name")
	if err != nil {
		log.Fatal("检查键存在失败:", err)
	}
	fmt.Println("删除后键是否存在:", exists)
}

// ExampleMemory 使用内存缓存示例
func ExampleMemory() {
	// 创建内存缓存实例
	cache := NewMemoryCache()
	defer cache.Close()

	// 设置键值对，过期时间10秒
	err := cache.Set("name", "go-cache-memory", 10*time.Second)
	if err != nil {
		log.Fatal("设置缓存失败:", err)
	}

	// 获取键值
	value, err := cache.Get("name")
	if err != nil {
		log.Fatal("获取缓存失败:", err)
	}
	fmt.Println("获取到的值:", value)

	// 检查键是否存在
	exists, err := cache.Exists("name")
	if err != nil {
		log.Fatal("检查键存在失败:", err)
	}
	fmt.Println("键是否存在:", exists)

	// 获取键的剩余生存时间
	ttl, err := cache.TTL("name")
	if err != nil {
		log.Fatal("获取TTL失败:", err)
	}
	fmt.Printf("剩余生存时间: %v\n", ttl)

	// 删除键
	err = cache.Delete("name")
	if err != nil {
		log.Fatal("删除键失败:", err)
	}

	// 再次检查键是否存在
	exists, err = cache.Exists("name")
	if err != nil {
		log.Fatal("检查键存在失败:", err)
	}
	fmt.Println("删除后键是否存在:", exists)
}

// ExampleFile 使用文件缓存示例
func ExampleFile() {
	// 创建文件缓存实例
	cache, err := NewFileCache("./cache")
	if err != nil {
		log.Fatal("创建文件缓存失败:", err)
	}
	defer cache.Close()

	// 设置键值对，过期时间10秒
	err = cache.Set("name", "go-cache-file", 10*time.Second)
	if err != nil {
		log.Fatal("设置缓存失败:", err)
	}

	// 获取键值
	value, err := cache.Get("name")
	if err != nil {
		log.Fatal("获取缓存失败:", err)
	}
	fmt.Println("获取到的值:", value)

	// 检查键是否存在
	exists, err := cache.Exists("name")
	if err != nil {
		log.Fatal("检查键存在失败:", err)
	}
	fmt.Println("键是否存在:", exists)

	// 获取键的剩余生存时间
	ttl, err := cache.TTL("name")
	if err != nil {
		log.Fatal("获取TTL失败:", err)
	}
	fmt.Printf("剩余生存时间: %v\n", ttl)

	// 删除键
	err = cache.Delete("name")
	if err != nil {
		log.Fatal("删除键失败:", err)
	}

	// 再次检查键是否存在
	exists, err = cache.Exists("name")
	if err != nil {
		log.Fatal("检查键存在失败:", err)
	}
	fmt.Println("删除后键是否存在:", exists)
}

// ExampleMulti 使用组合缓存示例
func ExampleMulti() {
	// 创建组合缓存实例
	memoryCache := NewMemoryCache()
	fileCache, err := NewFileCache("./cache")
	if err != nil {
		log.Fatal("创建文件缓存失败:", err)
	}

	multiCache := NewMultiCache(memoryCache, fileCache)
	defer multiCache.Close()

	// 设置键值对，过期时间10秒
	err = multiCache.Set("name", "go-cache-multi", 10*time.Second)
	if err != nil {
		log.Fatal("设置缓存失败:", err)
	}

	// 获取键值
	value, err := multiCache.Get("name")
	if err != nil {
		log.Fatal("获取缓存失败:", err)
	}
	fmt.Println("获取到的值:", value)

	// 检查键是否存在
	exists, err := multiCache.Exists("name")
	if err != nil {
		log.Fatal("检查键存在失败:", err)
	}
	fmt.Println("键是否存在:", exists)

	// 获取键的剩余生存时间
	ttl, err := multiCache.TTL("name")
	if err != nil {
		log.Fatal("获取TTL失败:", err)
	}
	fmt.Printf("剩余生存时间: %v\n", ttl)

	// 删除键
	err = multiCache.Delete("name")
	if err != nil {
		log.Fatal("删除键失败:", err)
	}

	// 再次检查键是否存在
	exists, err = multiCache.Exists("name")
	if err != nil {
		log.Fatal("检查键存在失败:", err)
	}
	fmt.Println("删除后键是否存在:", exists)
}
