package memdriver

type netAddr struct {
	h string
}

func (na *netAddr) Network() string {
	return "tcp"
}

func (na *netAddr) String() string {
	return na.h
}
