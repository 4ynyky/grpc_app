package memdriver

import (
	"bufio"
	"fmt"
)

func (mcd *memCacheDriver) Get(key string) (string, error) {
	conn, err := mcd.connPool.GetConn()
	if err != nil {
		return "", fmt.Errorf("Get connection failed: %w", err)
	}
	defer mcd.connPool.PutConn(conn)
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	if _, err = fmt.Fprintf(rw, "gets %s\r\n", key); err != nil {
		return "", err
	}
	if err = rw.Flush(); err != nil {
		return "", err
	}
	item, err := readItem(rw.Reader)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

func (mcd *memCacheDriver) Set(item Item) error {
	conn, err := mcd.connPool.GetConn()
	if err != nil {
		return fmt.Errorf("Get connection failed: %w", err)
	}
	defer mcd.connPool.PutConn(conn)
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	err = writeItem(rw, "set", item)
	if err != nil {
		return fmt.Errorf("failed write item: %w", err)
	}
	return nil
}

func (mcd *memCacheDriver) Delete(key string) error {
	conn, err := mcd.connPool.GetConn()
	if err != nil {
		return fmt.Errorf("Get connection failed: %w", err)
	}
	defer mcd.connPool.PutConn(conn)
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	err = deleteItem(rw, key)
	if err != nil {
		return fmt.Errorf("failed delete item: %w", err)
	}
	return nil
}
