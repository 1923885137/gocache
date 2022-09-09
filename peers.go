package gocache

import "GoCache/pd"

//PeerPicker is the interface that must be  implemented to locate
//the peer that owns a specific key
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

//PeerGetter is the interface that must be implemented by a peer
type PeerGetter interface {
	Get(in *pd.Request, out *pd.Response) error
}
