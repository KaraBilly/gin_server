package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	c = cache.New(1*time.Minute, 1*time.Minute)
)

//func init(){
//	c = cache.New(1*time.Minute, 1*time.Minute)
//}

func GetTokenCache(key string) (interface{}, bool) {
	return c.Get(AccessToken)
}
func SetTokenCache(key string, value string, expire time.Duration) {
	c.Set(key, value, expire)
}
