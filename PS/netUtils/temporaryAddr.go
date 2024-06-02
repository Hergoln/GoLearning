package netUtils

type MyAddr struct {
	network string
	addr string
}

func (addr MyAddr) Network() string {
	return addr.network
}

func (addr MyAddr) String() string {
	return addr.addr
}