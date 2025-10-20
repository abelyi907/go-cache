package go_cache

var redisServer = new(RedisCache)
var redisUrl = "localhost:36379"
var redisPassword = ""
var redisDb = 0

func init() {
	redisServer = NewRedisCache(redisUrl, redisPassword, redisDb)
}
