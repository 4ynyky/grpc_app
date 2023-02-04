package connectionpool

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type ConnectionPool struct {
	mu sync.Mutex

	addr         net.Addr
	conns        []net.Conn
	maxIdleConns int
	connTimeout  time.Duration
}

func New(addr net.Addr, connectionTimeout time.Duration, maxIdleConnections int) (*ConnectionPool, error) {
	cp := &ConnectionPool{}
	cp.addr = addr
	cp.connTimeout = connectionTimeout
	cp.maxIdleConns = maxIdleConnections
	return cp, nil
}

func (cp *ConnectionPool) GetConn() (net.Conn, error) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if cp.conns == nil {
		cp.conns = make([]net.Conn, 0)
	}

	if len(cp.conns) == 0 {
		conn, err := cp.openCon()
		if err != nil {
			return nil, fmt.Errorf("failed open connection: %w", err)
		}
		return conn, nil
	}

	conn := cp.conns[len(cp.conns)-1]
	cp.conns = cp.conns[:len(cp.conns)-1]
	return conn, nil
}

func (cp *ConnectionPool) PutConn(conn net.Conn) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if cp.conns == nil {
		cp.conns = make([]net.Conn, 0)
	}

	if len(cp.conns) >= cp.maxIdleConns {
		conn.Close()
		return
	}

	cp.conns = append(cp.conns, conn)
}

func (cp *ConnectionPool) openCon() (net.Conn, error) {
	nc, err := net.DialTimeout(cp.addr.Network(), cp.addr.String(), cp.connTimeout)
	if err == nil {
		return nc, nil
	}

	if ne, ok := err.(net.Error); ok && ne.Timeout() {
		return nil, ErrConnTimeout
	}

	return nil, err
}
