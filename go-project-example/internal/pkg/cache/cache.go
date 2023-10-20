package cache

import "go-project-example/internal/pkg/cache/redis"

func Initialize(host, password, prefix string, database int) (err error) {
	return redis.Initialize(host, password, prefix, database)
}

func GetRedis() *redis.Redis {
	return redis.Get()
}
