package gossipdb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMemberListWhenThereIsOnlyOneMember(t *testing.T) {
	memberList, err := newMemberlist(9000, "", nil)
	node := memberList.LocalNode()
	assert.Nil(t, err)
	assert.Equal(t, 1, memberList.NumMembers())
	assert.Equal(t, uint16(9000), node.Port)
}

func TestNewMemberListWhenThereAreMoreThanOneMembers(t *testing.T) {
	memberList, err := newMemberlist(9001, "0.0.0.0:9000,0.0.0.0:9001", nil)
	node := memberList.LocalNode()
	assert.Nil(t, err)
	assert.Equal(t, 2, memberList.NumMembers())
	assert.Equal(t, uint16(9001), node.Port)
}

func TestNewMemberListWhenAddressIsAlreadyInUse(t *testing.T) {
	_, err := newMemberlist(9001, "", nil)
	assert.NotNil(t, err)
}
