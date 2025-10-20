package go_cache

import (
	"fmt"
	"log"
	"time"

	go_cache "github.com/yourusername/go-cache"
)

func main() {
	fmt.Println("=== Redis缓存示例 ===")
	redisExample()

	fmt.Println("\n=== 内存缓存示例 ===")
	memoryExample()

	fmt.Println("\n=== 文件缓存示例 ===")
	fileExample()

	fmt.Println("\n=== 组合缓存示例 ===")
	multiExample()
}

func redisExample() {
	// 创建Redis缓存实例
	// 注意：需要确保Redis服务器正在运行，并且地址正确
	cache := go_cache.NewRedisCache("localhost:6379", "", 0)
	defer func() {
		if err := cache.Close(); err != nil {
			log.Printf("关闭Redis连接时出错: %v", err)
		}
	}()

	// 设置键值对，过期时间10秒
	fmt.Println("设置键值对...")
	err := cache.Set("name", "go-cache-redis", 10*time.Second)
	if err != nil {
		log.Fatalf("设置缓存失败: %v", err)
	}

	// 获取键值
	fmt.Println("获取键值...")
	value, err := cache.Get("name")
	if err != nil {
		log.Fatalf("获取缓存失败: %v", err)
	}
	fmt.Printf("获取到的值: %s\n", value)

	// 检查键是否存在
	fmt.Println("检查键是否存在...")
	exists, err := cache.Exists("name")
	if err != nil {
		log.Fatalf("检查键存在失败: %v", err)
	}
	fmt.Printf("键是否存在: %t\n", exists)

	// 获取键的剩余生存时间
	fmt.Println("获取键的剩余生存时间...")
	ttl, err := cache.TTL("name")
	if err != nil {
		log.Fatalf("获取TTL失败: %v", err)
	}
	fmt.Printf("剩余生存时间: %v\n", ttl)

	// 设置新的过期时间
	fmt.Println("设置新的过期时间...")
	err = cache.Expire("name", 20*time.Second)
	if err != nil {
		log.Fatalf("设置过期时间失败: %v", err)
	}

	// 再次获取TTL
	ttl, err = cache.TTL("name")
	if err != nil {
		log.Fatalf("获取TTL失败: %v", err)
	}
	fmt.Printf("新的剩余生存时间: %v\n", ttl)

	// 删除键
	fmt.Println("删除键...")
	err = cache.Delete("name")
	if err != nil {
		log.Fatalf("删除键失败: %v", err)
	}

	// 再次检查键是否存在
	fmt.Println("删除后再次检查键是否存在...")
	exists, err = cache.Exists("name")
	if err != nil {
		log.Fatalf("检查键存在失败: %v", err)
	}
	fmt.Printf("删除后键是否存在: %t\n", exists)
}

func memoryExample() {
	// 创建内存缓存实例
	cache := go_cache.NewMemoryCache()
	defer func() {
		if err := cache.Close(); err != nil {
			log.Printf("关闭内存缓存时出错: %v", err)
		}
	}()

	// 设置键值对，过期时间10秒
	fmt.Println("设置键值对...")
	err := cache.Set("name", "go-cache-memory", 10*time.Second)
	if err != nil {
		log.Fatalf("设置缓存失败: %v", err)
	}

	// 获取键值
	fmt.Println("获取键值...")
	value, err := cache.Get("name")
	if err != nil {
		log.Fatalf("获取缓存失败: %v", err)
	}
	fmt.Printf("获取到的值: %s\n", value)

	// 检查键是否存在
	fmt.Println("检查键是否存在...")
	exists, err := cache.Exists("name")
	if err != nil {
		log.Fatalf("检查键存在失败: %v", err)
	}
	fmt.Printf("键是否存在: %t\n", exists)

	// 获取键的剩余生存时间
	fmt.Println("获取键的剩余生存时间...")
	ttl, err := cache.TTL("name")
	if err != nil {
		log.Fatalf("获取TTL失败: %v", err)
	}
	fmt.Printf("剩余生存时间: %v\n", ttl)

	// 设置新的过期时间
	fmt.Println("设置新的过期时间...")
	err = cache.Expire("name", 20*time.Second)
	if err != nil {
		log.Fatalf("设置过期时间失败: %v", err)
	}

	// 再次获取TTL
	ttl, err = cache.TTL("name")
	if err != nil {
		log.Fatalf("获取TTL失败: %v", err)
	}
	fmt.Printf("新的剩余生存时间: %v\n", ttl)

	// 删除键
	fmt.Println("删除键...")
	err = cache.Delete("name")
	if err != nil {
		log.Fatalf("删除键失败: %v", err)
	}

	// 再次检查键是否存在
	fmt.Println("删除后再次检查键是否存在...")
	exists, err = cache.Exists("name")
	if err != nil {
		log.Fatalf("检查键存在失败: %v", err)
	}
	fmt.Printf("删除后键是否存在: %t\n", exists)
}

func fileExample() {
	// 创建文件缓存实例
	cache, err := go_cache.NewFileCache("./cache")
	if err != nil {
		log.Fatalf("创建文件缓存失败: %v", err)
	}
	defer func() {
		if err := cache.Close(); err != nil {
			log.Printf("关闭文件缓存时出错: %v", err)
		}
	}()

	// 设置键值对，过期时间10秒
	fmt.Println("设置键值对...")
	err = cache.Set("name", "go-cache-file", 10*time.Second)
	if err != nil {
		log.Fatalf("设置缓存失败: %v", err)
	}

	// 获取键值
	fmt.Println("获取键值...")
	value, err := cache.Get("name")
	if err != nil {
		log.Fatalf("获取缓存失败: %v", err)
	}
	fmt.Printf("获取到的值: %s\n", value)

	// 检查键是否存在
	fmt.Println("检查键是否存在...")
	exists, err := cache.Exists("name")
	if err != nil {
		log.Fatalf("检查键存在失败: %v", err)
	}
	fmt.Printf("键是否存在: %t\n", exists)

	// 获取键的剩余生存时间
	fmt.Println("获取键的剩余生存时间...")
	ttl, err := cache.TTL("name")
	if err != nil {
		log.Fatalf("获取TTL失败: %v", err)
	}
	fmt.Printf("剩余生存时间: %v\n", ttl)

	// 设置新的过期时间
	fmt.Println("设置新的过期时间...")
	err = cache.Expire("name", 20*time.Second)
	if err != nil {
		log.Fatalf("设置过期时间失败: %v", err)
	}

	// 再次获取TTL
	ttl, err = cache.TTL("name")
	if err != nil {
		log.Fatalf("获取TTL失败: %v", err)
	}
	fmt.Printf("新的剩余生存时间: %v\n", ttl)

	// 删除键
	fmt.Println("删除键...")
	err = cache.Delete("name")
	if err != nil {
		log.Fatalf("删除键失败: %v", err)
	}

	// 再次检查键是否存在
	fmt.Println("删除后再次检查键是否存在...")
	exists, err = cache.Exists("name")
	if err != nil {
		log.Fatalf("检查键存在失败: %v", err)
	}
	fmt.Printf("删除后键是否存在: %t\n", exists)
}

func multiExample() {
	// 创建组合缓存实例
	memoryCache := go_cache.NewMemoryCache()
	fileCache, err := go_cache.NewFileCache("./cache")
	if err != nil {
		log.Fatalf("创建文件缓存失败: %v", err)
	}

	multiCache := go_cache.NewMultiCache(memoryCache, fileCache)
	defer func() {
		if err := multiCache.Close(); err != nil {
			log.Printf("关闭组合缓存时出错: %v", err)
		}
	}()

	// 设置键值对，过期时间10秒
	fmt.Println("设置键值对...")
	err = multiCache.Set("name", "go-cache-multi", 10*time.Second)
	if err != nil {
		log.Fatalf("设置缓存失败: %v", err)
	}

	// 获取键值
	fmt.Println("获取键值...")
	value, err := multiCache.Get("name")
	if err != nil {
		log.Fatalf("获取缓存失败: %v", err)
	}
	fmt.Printf("获取到的值: %s\n", value)

	// 检查键是否存在
	fmt.Println("检查键是否存在...")
	exists, err := multiCache.Exists("name")
	if err != nil {
		log.Fatalf("检查键存在失败: %v", err)
	}
	fmt.Printf("键是否存在: %t\n", exists)

	// 获取键的剩余生存时间
	fmt.Println("获取键的剩余生存时间...")
	ttl, err := multiCache.TTL("name")
	if err != nil {
		log.Fatalf("获取TTL失败: %v", err)
	}
	fmt.Printf("剩余生存时间: %v\n", ttl)

	// 设置新的过期时间
	fmt.Println("设置新的过期时间...")
	err = multiCache.Expire("name", 20*time.Second)
	if err != nil {
		log.Fatalf("设置过期时间失败: %v", err)
	}

	// 再次获取TTL
	ttl, err = multiCache.TTL("name")
	if err != nil {
		log.Fatalf("获取TTL失败: %v", err)
	}
	fmt.Printf("新的剩余生存时间: %v\n", ttl)

	// 删除键
	fmt.Println("删除键...")
	err = multiCache.Delete("name")
	if err != nil {
		log.Fatalf("删除键失败: %v", err)
	}

	// 再次检查键是否存在
	fmt.Println("删除后再次检查键是否存在...")
	exists, err = multiCache.Exists("name")
	if err != nil {
		log.Fatalf("检查键存在失败: %v", err)
	}
	fmt.Printf("删除后键是否存在: %t\n", exists)
}
