package redis

import (
	"github.com/go-redis/redis"
	"time"
)

var rc *Redis

type Redis struct {
	Client *redis.Client
	Prefix string
}

func Initialize(host, password, prefix string, database int) (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       database,
	})
	pong, err := client.Ping().Result()
	if pong != "PONG" || err != nil {
		return
	}

	rc = &Redis{
		Client: client,
		Prefix: prefix,
	}

	return
}

func Get() *Redis {
	return rc
}

func (r *Redis) Get(key string) (value string, err error) {
	return r.Client.Get(r.Prefix + key).Result()
}

func (r *Redis) MGet(keys ...string) (values []string, err error) {
	for i, key := range keys {
		keys[i] = r.Prefix + key
	}

	var results []interface{}
	if results, err = r.Client.MGet(keys...).Result(); err != nil {
		return
	}

	values = make([]string, len(keys))
	for i := range results {
		values[i] = results[i].(string)
	}

	return
}

func (r *Redis) Set(key, value string, expire time.Duration) (err error) {
	return r.Client.Set(r.Prefix + key, value, expire).Err()
}

func (r *Redis) SetNX(key, value string, expire time.Duration) (ok bool, err error) {
	return r.Client.SetNX(r.Prefix + key, value, expire).Result()
}

func (r *Redis) MSet(values map[string]string, expire time.Duration) (err error) {
	var pipe = r.Client.Pipeline()
	for key, value := range values {
		pipe.Set(r.Prefix + key, value, expire)
	}
	if _, err = pipe.Exec(); err != nil {
		return
	}

	return
}

func (r *Redis) Del(key string) (err error) {
	return r.Client.Del(r.Prefix + key).Err()
}

func (r *Redis) Close() (err error) {
	return r.Client.Close()
}