package gossipdb

import (
	"github.com/patrickmn/go-cache"
)

type db struct {
	connection *cache.Cache
}

func newDb() *db {
	return &db{
		connection: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (d *db) Save(k string, b []byte) {
	d.connection.Set(k, b, cache.DefaultExpiration)
}

func (d *db) Get(k string) ([]byte, bool) {
	v, f := d.connection.Get(k)
	return v.([]byte), f
}
