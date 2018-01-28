package gossipdb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testGossipDbWithMoreThanOneNode(t *testing.T) {
	gossipDbOne, errOne := NewGossipDb("0.0.0.0:9000, 0.0.0.0:9001", 9000)
	gossipDbTwo, errTwo := NewGossipDb("0.0.0.0:9000, 0.0.0.0:9001", 9001)
	assert.Nil(t, errOne)
	assert.Nil(t, errTwo)
	assert.Equal(t, gossipDbOne.Members(), gossipDbTwo.Members())
	assert.Equal(t, gossipDbOne.members.LocalNode().Port, 9000)
	assert.Equal(t, gossipDbTwo.members.LocalNode().Port, 9001)
	assert.Equal(t, 2, gossipDbOne.members.NumMembers())
	assert.Equal(t, 2, gossipDbOne.broadcasts.NumNodes())
	assert.Equal(t, 3, gossipDbOne.broadcasts.RetransmitMult)
	defer gossipDbOne.Shutdown()
	defer gossipDbTwo.Shutdown()
}

func TestNewGossipDbWhenMemberAddressIsNotInUse(t *testing.T) {
	gossipDb, err := NewGossipDb("", 9000)
	assert.Nil(t, err)
	assert.Equal(t, 1, gossipDb.members.NumMembers())
	assert.Equal(t, 1, gossipDb.broadcasts.NumNodes())
	assert.NotNil(t, gossipDb.database.connection)
	defer gossipDb.Shutdown()
}

func TestGossipDbWhenMemberAddressIsAlreadyInUse(t *testing.T) {
	gossipDb, errOne := NewGossipDb("", 9000)
	_, errTwo := NewGossipDb("", 9000)
	assert.Nil(t, errOne)
	assert.NotNil(t, errTwo)
	defer gossipDb.Shutdown()
}

func TestGetWhenKeyIsNotFound(t *testing.T) {
	gossipDb, _ := NewGossipDb("", 9000)
	value, found := gossipDb.Get("ping")
	assert.False(t, found)
	assert.Equal(t, value, "nil")
	defer gossipDb.Shutdown()
}

func TestGetWhenKeyIsFound(t *testing.T) {
	gossipDb, _ := NewGossipDb("", 9000)
	gossipDb.Set("ping", "pong")
	value, found := gossipDb.Get("ping")
	assert.True(t, found)
	assert.Equal(t, "pong", value)
	defer gossipDb.Shutdown()
}

func TestSetWhenKeyAlreadyExist(t *testing.T) {
	gossipDb, _ := NewGossipDb("", 9000)
	gossipDb.Set("ping", "pong")
	value, _ := gossipDb.Get("ping")
	assert.Equal(t, "pong", value)
	gossipDb.Set("ping", "pong2")
	newValue, _ := gossipDb.Get("ping")
	assert.Equal(t, "pong2", newValue)
	defer gossipDb.Shutdown()
}

func TestMembers(t *testing.T) {
	gossipDbOne, _ := NewGossipDb("0.0.0.0:9000,0.0.0.0:9001", 9000)
	gossipDbTwo, _ := NewGossipDb("0.0.0.0:9000,0.0.0.0:9001", 9001)
	assert.Equal(t, 2, len(gossipDbOne.Members()))
	gossipDbTwo.Shutdown()
	gossipDbOne.Shutdown()
}
