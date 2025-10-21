package go_cache

import (
	"os"
)

var redisServer = new(RedisCache)
var redisUrl = "localhost:36379"
var redisPassword = ""
var redisDb = 0
var testFilePath = "./test_cache"

func init() {
	redisServer = NewRedisCache(redisUrl, redisPassword, redisDb, "gocache:")
}
func Init() func() {
	clearTestFile()
	return func() {
		clearTestFile()
	}
}

// 每次测试结束后，清理测试文件
func clearTestFile() {
	err := os.RemoveAll(testFilePath)
	if err != nil {
		panic(err)
	}
}
