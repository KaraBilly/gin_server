package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	c = cache.New(1*time.Minute, 1*time.Minute)
)

func GetTokenCache(key string) (string, bool) {
	token, exist := c.Get(AccessToken)
	return token.(string), exist
}
func SetTokenCache(key string, value string, expire time.Duration) {
	c.Set(key, value, expire)
}
