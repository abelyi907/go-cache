# 这是一个支持多种存储后端的缓存库

## 简介

这是一个支持多种存储后端的Go语言缓存库，目前支持：
- Redis缓存
- 内存缓存
- 文件系统缓存

提供了统一的缓存接口，可以单独使用任意一种缓存，也可以组合使用多种缓存。

## 功能特性

- 支持Redis、内存、文件系统三种缓存后端
- 统一的缓存接口，便于切换和组合使用
- 支持设置键值对并指定过期时间
- 支持获取指定键的值
- 支持删除指定键
- 支持检查键是否存在
- 支持设置键的过期时间
- 支持获取键的剩余生存时间
- 支持组合多种缓存后端的MultiCache

## 安装

确保你已经安装了Go环境（推荐1.16及以上版本）。

```bash
go mod tidy
```

## 依赖

本项目依赖于 [go-redis/redis](https://github.com/go-redis/redis) v8 版本。

## 使用方法

首先确保你有一个正在运行的Redis服务器（如果使用Redis缓存）。

### 使用单一缓存

#### Redis缓存

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    go_cache "github.com/yourusername/go-cache"
)

func main() {
    // 创建Redis缓存实例
    cache := go_cache.NewRedisCache("localhost:6379", "", 0)
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
}
```

#### 内存缓存

```go
// 创建内存缓存实例
cache := go_cache.NewMemoryCache()
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
```

#### 文件缓存

```go
// 创建文件缓存实例
cache, err := go_cache.NewFileCache("./cache")
if err != nil {
    log.Fatal("创建文件缓存失败:", err)
}
defer cache.Close()

// 设置键值对，过期时间10秒
err = cache.Set("name", "go-cache", 10*time.Second)
if err != nil {
    log.Fatal("设置缓存失败:", err)
}

// 获取键值
value, err := cache.Get("name")
if err != nil {
    log.Fatal("获取缓存失败:", err)
}
fmt.Println("获取到的值:", value)
```

### 使用工厂方法创建缓存

```go
// 创建Redis缓存
redisConfig := go_cache.CacheConfig{
    Type:          go_cache.RedisCacheType,
    RedisAddr:     "localhost:6379",
    RedisPassword: "",
    RedisDB:       0,
}
redisCache, _ := go_cache.NewCache(redisConfig)

// 创建内存缓存
memoryConfig := go_cache.CacheConfig{
    Type: go_cache.MemoryCacheType,
}
memoryCache := go_cache.NewCache(memoryConfig)

// 创建文件缓存
fileConfig := go_cache.CacheConfig{
    Type:    go_cache.FileCacheType,
    FileDir: "./cache",
}
fileCache, _ := go_cache.NewCache(fileConfig)
```

### 使用组合缓存（MultiCache）

```go
// 创建组合缓存，同时使用内存和Redis缓存
memoryCache := go_cache.NewMemoryCache()
redisCache := go_cache.NewRedisCache("localhost:6379", "", 0)
multiCache := go_cache.NewMultiCache(memoryCache, redisCache)
defer multiCache.Close()

// 设置键值对
err := multiCache.Set("name", "go-cache", 10*time.Second)
if err != nil {
    log.Fatal("设置缓存失败:", err)
}

// 获取键值（会先从内存查找，找不到再从Redis查找，并将结果回填到内存中）
value, err := multiCache.Get("name")
if err != nil {
    log.Fatal("获取缓存失败:", err)
}
fmt.Println("获取到的值:", value)
```

## API参考

### Cache接口

所有缓存实现都遵循统一的Cache接口：

#### Set(key string, value interface{}, expiration time.Duration) error

将键值对存储到缓存中，并设置过期时间。

#### Get(key string) (string, error)

从缓存中获取指定键的值。

#### Delete(key string) error

从缓存中删除指定键。

#### Exists(key string) (bool, error)

检查指定键是否存在于缓存中。

#### Expire(key string, expiration time.Duration) error

设置键的过期时间。

#### TTL(key string) (time.Duration, error)

获取键的剩余生存时间。

#### Close() error

关闭缓存连接。

### 工厂方法

#### NewCache(config CacheConfig) (Cache, error)

根据配置创建缓存实例。

### 特定实现构造函数

#### NewRedisCache(addr, password string, db int) *RedisCache

创建Redis缓存实例。

#### NewMemoryCache() *MemoryCache

创建内存缓存实例。

#### NewFileCache(dir string) (*FileCache, error)

创建文件缓存实例。

#### NewMultiCache(caches ...Cache) *MultiCache

创建组合缓存实例。

## 运行示例

```bash
go run main.go
```

## 运行测试

```bash
go test -v
```