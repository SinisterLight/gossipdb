package gossipdb

import (
	"github.com/hashicorp/memberlist"
)

type KeyParser interface {
	ToKey(b []byte) string
}

type GossipDb struct {
	members    *memberlist.Memberlist
	broadcasts *memberlist.TransmitLimitedQueue
	database   *db
	keyParser  KeyParser
}

func NewGossipDb(members string, port int, k KeyParser) (*GossipDb, error) {
	d := newDb()

	b := &memberlist.TransmitLimitedQueue{
		RetransmitMult: 3,
	}

	del := &delegate{
		getBroadcasts: func(overhead, limit int) [][]byte {
			return b.GetBroadcasts(overhead, limit)
		},
		notifyMsg: func(b []byte) {
			k := k.ToKey(b)
			d.Save(k, b)
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
		keyParser:  k,
	}, nil
}

func (gdb *GossipDb) Get(k string) ([]byte, bool) {
	return gdb.database.Get(k)
}

func (gdb *GossipDb) Set(m []byte) {
	k := gdb.keyParser.ToKey(m)
	gdb.database.Save(k, m)
	gdb.broadcasts.QueueBroadcast(&broadcast{
		msg:    m,
		notify: nil,
	})
}
