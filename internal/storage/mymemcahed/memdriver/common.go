package memdriver

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func readItem(r *bufio.Reader) (*Item, error) {
	line, err := r.ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	if bytes.Equal(line, resultEnd) {
		return nil, nil
	}
	it := new(Item)
	size, err := scanLine(line, it)
	if err != nil {
		return nil, err
	}
	offset := 2
	it.Value = make([]byte, size+offset)
	_, err = io.ReadFull(r, it.Value)
	if err != nil {
		it.Value = nil
		return nil, err
	}
	if !bytes.HasSuffix(it.Value, crlf) {
		it.Value = nil
		return nil, fmt.Errorf("memcache: corrupt get result read")
	}
	it.Value = it.Value[:size]
	return it, nil
}

func scanLine(line []byte, it *Item) (size int, err error) {
	/*
		VALUE <key> <flags> <bytes> [<cas unique>]\r\n
		<data block>\r\n

		- <key> is the key for the item being sent

		- <flags> is the flags value set by the storage command

		- <bytes> is the length of the data block to follow, *not* including
		its delimiting \r\n

		- <cas unique> is a unique 64-bit integer that uniquely identifies
		this specific item.

		- <data block> is the data for this item.
	*/
	spacesCount := 3
	pattern := "VALUE %s %d %d %d\r\n"
	dest := []interface{}{&it.Key, &it.Flags, &size, &it.casid}
	if bytes.Count(line, space) == spacesCount {
		pattern = "VALUE %s %d %d\r\n"
		dest = dest[:3]
	}
	n, err := fmt.Sscanf(string(line), pattern, dest...)
	if err != nil || n != len(dest) {
		return -1, fmt.Errorf("memcache: unexpected line in get response: %q", line)
	}
	return size, nil
}

func writeItem(rw *bufio.ReadWriter, verb string, item Item) error {
	if !checkKey(item.Key) {
		return ErrMalformedKey
	}

	var err error

	_, err = fmt.Fprintf(rw, "%s %s %d %d %d\r\n",
		verb, item.Key, item.Flags, item.Expiration, len(item.Value))
	if err != nil {
		return err
	}
	if _, err = rw.Write(item.Value); err != nil {
		return err
	}
	if _, err = rw.Write(crlf); err != nil {
		return err
	}
	if err = rw.Flush(); err != nil {
		return err
	}
	line, err := rw.ReadSlice('\n')
	if err != nil {
		return err
	}
	switch {
	case bytes.Equal(line, resultStored):
		return nil
	case bytes.Equal(line, resultNotStored):
		return ErrNotStored
	case bytes.Equal(line, resultExists):
		return ErrCASConflict
	case bytes.Equal(line, resultNotFound):
		return ErrCacheMiss
	}
	return fmt.Errorf("memcache: unexpected response line from %q: %q", verb, string(line))
}

func checkKey(key string) bool {
	if len(key) > maxKeyLen {
		return false
	}
	for i := 0; i < len(key); i++ {
		if key[i] <= ' ' || key[i] == 0x7f {
			return false
		}
	}
	return true
}

func deleteItem(rw *bufio.ReadWriter, key string) error {
	_, err := fmt.Fprintf(rw, "delete %s\r\n", key)
	if err != nil {
		return err
	}
	if err = rw.Flush(); err != nil {
		return err
	}
	line, err := rw.ReadSlice('\n')
	if err != nil {
		return err
	}
	switch {
	case bytes.Equal(line, resultOK):
		return nil
	case bytes.Equal(line, resultDeleted):
		return nil
	case bytes.Equal(line, resultNotStored):
		return ErrNotStored
	case bytes.Equal(line, resultExists):
		return ErrCASConflict
	case bytes.Equal(line, resultNotFound):
		return ErrCacheMiss
	}
	return fmt.Errorf("memcache: unexpected response line: %q", string(line))
}
