package go_cache

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache 实现了基于Redis的缓存库
type RedisCache struct {
	client    *redis.Client
	ctx       context.Context
	prefixKey string
}

// NewRedisCache 创建一个新的Redis缓存实例
func NewRedisCache(addr string, password string, db int, PrefixKey string) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{
		client:    client,
		ctx:       context.Background(),
		prefixKey: PrefixKey,
	}
}

// Set 将键值对存储到缓存中，并设置过期时间
func (r *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(r.ctx, r.prefixKey+key, value, expiration).Err()
}

// Get 从缓存中获取指定键的值
func (r *RedisCache) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, r.prefixKey+key).Result()
	if errors.Is(err, redis.Nil) {
		return "", ErrKeyNotFound
	}
	return val, err
}

// Delete 从缓存中删除指定键
func (r *RedisCache) Delete(key string) error {
	return r.client.Del(r.ctx, r.prefixKey+key).Err()
}

// Exists 检查指定键是否存在于缓存中
func (r *RedisCache) Exists(key string) (bool, error) {
	result, err := r.client.Exists(r.ctx, r.prefixKey+key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

// Expire 设置键的过期时间
func (r *RedisCache) Expire(key string, expiration time.Duration) error {
	result, err := r.client.Expire(r.ctx, r.prefixKey+key, expiration).Result()
	if err != nil {
		return err
	}
	if !result {
		return ErrKeyNotFound
	}
	return nil
}

// TTL 获取键的剩余生存时间
func (r *RedisCache) TTL(key string) (time.Duration, error) {
	ttl, err := r.client.TTL(r.ctx, r.prefixKey+key).Result()
	if err != nil {
		return 0, err
	}

	if ttl == time.Duration(-1) {
		// Redis中-1表示永不过期
		return time.Duration(-1), nil
	}

	if ttl == time.Duration(-2) {
		// Redis中-2表示键不存在
		return time.Duration(-2), ErrKeyNotFound
	}

	return ttl, nil
}

// Close 关闭Redis连接
func (r *RedisCache) Close() error {
	return r.client.Close()
}
