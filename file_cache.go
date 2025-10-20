package go_cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// FileCache 实现了基于文件系统的缓存
type FileCache struct {
	dir string
}

// fileItem 表示文件缓存中的一个项目
type fileItem struct {
	Value      string    `json:"value"`
	Expiration time.Time `json:"expiration"`
}

// NewFileCache 创建一个新的文件系统缓存实例
func NewFileCache(dir string) (*FileCache, error) {
	// 确保目录存在
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	return &FileCache{
		dir: dir,
	}, nil
}

// Set 将键值对存储到缓存中，并设置过期时间
func (f *FileCache) Set(key string, value interface{}, expiration time.Duration) error {
	var expirationTime time.Time
	if expiration > 0 {
		expirationTime = time.Now().Add(expiration)
	}

	item := &fileItem{
		Value:      value.(string),
		Expiration: expirationTime,
	}

	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	filePath := f.getFilePath(key)
	return os.WriteFile(filePath, data, 0644)
}

// Get 从缓存中获取指定键的值
func (f *FileCache) Get(key string) (string, error) {
	filePath := f.getFilePath(key)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", ErrKeyNotFound
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	var item fileItem
	err = json.Unmarshal(data, &item)
	if err != nil {
		return "", err
	}

	// 检查是否过期
	if !item.Expiration.IsZero() && time.Now().After(item.Expiration) {
		// 删除过期文件
		os.Remove(filePath)
		return "", ErrKeyNotFound
	}

	return item.Value, nil
}

// Delete 从缓存中删除指定键
func (f *FileCache) Delete(key string) error {
	filePath := f.getFilePath(key)
	
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil // 文件不存在，认为删除成功
	}
	
	return os.Remove(filePath)
}

// Exists 检查指定键是否存在于缓存中
func (f *FileCache) Exists(key string) (bool, error) {
	filePath := f.getFilePath(key)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	var item fileItem
	err = json.Unmarshal(data, &item)
	if err != nil {
		// 数据损坏，删除文件
		os.Remove(filePath)
		return false, nil
	}

	// 检查是否过期
	if !item.Expiration.IsZero() && time.Now().After(item.Expiration) {
		// 删除过期文件
		os.Remove(filePath)
		return false, nil
	}

	return true, nil
}

// Expire 设置键的过期时间
func (f *FileCache) Expire(key string, expiration time.Duration) error {
	filePath := f.getFilePath(key)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return ErrKeyNotFound
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var item fileItem
	err = json.Unmarshal(data, &item)
	if err != nil {
		return err
	}

	if expiration > 0 {
		item.Expiration = time.Now().Add(expiration)
	} else {
		item.Expiration = time.Time{}
	}

	// 保存更新后的项
	newData, err := json.Marshal(item)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, newData, 0644)
}

// TTL 获取键的剩余生存时间
func (f *FileCache) TTL(key string) (time.Duration, error) {
	filePath := f.getFilePath(key)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return 0, ErrKeyNotFound
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	var item fileItem
	err = json.Unmarshal(data, &item)
	if err != nil {
		return 0, err
	}

	// 检查是否过期
	if !item.Expiration.IsZero() && time.Now().After(item.Expiration) {
		// 删除过期文件
		os.Remove(filePath)
		return 0, ErrKeyNotFound
	}

	if item.Expiration.IsZero() {
		// 永不过期
		return -1, nil
	}

	return time.Until(item.Expiration), nil
}

// Close 关闭缓存连接
func (f *FileCache) Close() error {
	// 文件系统缓存不需要特殊关闭操作
	return nil
}

// getFilePath 获取键对应的文件路径
func (f *FileCache) getFilePath(key string) string {
	// 简单的键到文件名的转换
	// 在实际应用中，可能需要更复杂的处理来避免文件名冲突
	filename := fmt.Sprintf("%x.json", key)
	return filepath.Join(f.dir, filename)
}