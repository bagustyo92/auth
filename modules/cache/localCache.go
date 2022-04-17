package cache

import (
	"errors"
	"fmt"
	"time"

	"github.com/bagustyo92/auth/config"
	"github.com/patrickmn/go-cache"
)

type localCache struct {
	db *cache.Cache
}

func DefaultLocalCache() Cacher {
	cache := cache.New(
		time.Duration(config.CacheDefaultTTL)*time.Second,
		time.Duration(config.CacheCleanInterval)*time.Second,
	)
	return &localCache{
		db: cache,
	}
}

func (l *localCache) Get(key interface{}) (interface{}, error) {
	value, ok := l.db.Get(fmt.Sprintf("%v", key))
	if !ok {
		return value, errors.New("No cache found")
	}

	return value, nil
}

func (l *localCache) Set(key interface{}, value interface{}, ttl int64) error {
	l.db.Set(fmt.Sprintf("%v", key), value, time.Duration(ttl)*time.Second)
	return nil
}

func (l *localCache) Delete(key interface{}) error {
	l.db.Delete(fmt.Sprintf("%v", key))
	return nil
}
