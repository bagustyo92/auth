package cache

type CacheConfigOptions struct {
	TTL int
}

type Cacher interface {
	Get(key interface{}) (interface{}, error)
	Set(key interface{}, value interface{}, ttl int64) error
	Delete(key interface{}) error
}
