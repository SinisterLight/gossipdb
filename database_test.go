package gossipdb

import (
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDb(t *testing.T) {
	expectedDb := &db{connection: cache.New(cache.NoExpiration, cache.NoExpiration)}
	actualDb := newDb()
	assert.Equal(t, expectedDb, actualDb)
}

func TestShouldReturnFalseIfKeyIsNotFoundInCache(t *testing.T) {
	db := newDb()
	_, found := db.Get("ping")
	assert.Equal(t, found, false)
}

func TestShouldReturnTrueIfKeyIsFoundInCache(t *testing.T) {
	db := newDb()
	db.Save("ping", "pong")
	value, found := db.Get("ping")
	assert.Equal(t, found, true)
	assert.Equal(t, value, "pong")
}
