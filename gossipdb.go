package gossipdb

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/memberlist"
)

type Pair struct {
	Key   string
	Value string
}

type GossipDb struct {
	members    *memberlist.Memberlist
	broadcasts *memberlist.TransmitLimitedQueue
	database   *db
}

func NewGossipDb(members string, port int) (*GossipDb, error) {
	d := newDb()

	b := &memberlist.TransmitLimitedQueue{
		RetransmitMult: 3,
	}

	del := &delegate{
		getBroadcasts: func(overhead, limit int) [][]byte {
			return b.GetBroadcasts(overhead, limit)
		},
		notifyMsg: func(b []byte) {
			pair := &Pair{}
			json.Unmarshal(b, pair)
			d.Save(pair.Key, pair.Value)
		},
	}

	m, err := newMemberlist(port, members, del)
	if err != nil {
		return nil, err
	}

	b.NumNodes = func() int {
		return m.NumMembers()
	}

	return &GossipDb{
		members:    m,
		broadcasts: b,
		database:   d,
	}, nil
}

func (gdb *GossipDb) Get(k string) (string, bool) {
	value, found := gdb.database.Get(k)
	if found {
		return value, found
	}
	return "nil", found
}

func (gdb *GossipDb) Set(key string, value string) {
	gdb.database.Save(key, value)
	pair := &Pair{Key: key, Value: value}
	message, err := json.Marshal(pair)
	if err != nil {
		fmt.Println(err)
	}

	gdb.broadcasts.QueueBroadcast(&broadcast{
		msg:    message,
		notify: nil,
	})
}

func (gdb *GossipDb) Members() []string {
	a := []string{}
	for _, m := range gdb.members.Members() {
		a = append(a, m.Name)
	}
	return a
}
