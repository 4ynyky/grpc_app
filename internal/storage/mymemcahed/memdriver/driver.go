package memdriver

import (
	"fmt"
	"net"
	"time"

	"github.com/4ynyky/grpc_app/internal/storage/mymemcahed/memdriver/connectionpool"
)

type connectionPooler interface {
	GetConn() (net.Conn, error)
	PutConn(conn net.Conn)
}

type memCacheDriver struct {
	connPool connectionPooler
}

func New(host string) (*memCacheDriver, error) {
	defTimeout := 1 * time.Second
	defMaxIdleConns := 4
	cPool, err := connectionpool.New(&netAddr{h: host}, defTimeout, defMaxIdleConns)
	if err != nil {
		return nil, fmt.Errorf("failed create connection pool: %w", err)
	}
	memCacheDriver := &memCacheDriver{
		connPool: cPool,
	}

	return memCacheDriver, nil
}
