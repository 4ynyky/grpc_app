package memcached

import (
	"fmt"

	"github.com/4ynyky/grpc_app/internal/domains"
	"github.com/4ynyky/grpc_app/internal/storage"
	"github.com/4ynyky/grpc_app/internal/storage/mymemcahed/memdriver"
)

type Config struct {
	Host string
}

type memDriver interface {
	Get(key string) (string, error)
	Set(item memdriver.Item) error
	Delete(key string) error
}

type memcachedStorage struct {
	md memDriver
}

func NewMemcachedStorage(conf Config) (*memcachedStorage, error) {
	var err error
	ms := &memcachedStorage{}
	ms.md, err = memdriver.New(conf.Host)
	if err != nil {
		// TODO for go1.20 raplace with join
		return nil, storage.ErrConnNotEstablished
	}
	return ms, nil
}

func (ms *memcachedStorage) Set(item domains.Item) error {
	if err := ms.md.Set(memdriver.Item{Key: item.ID, Value: []byte(item.Value)}); err != nil {
		return fmt.Errorf("failed store item: %v, error: %w", item, err)
	}
	return nil
}

func (ms *memcachedStorage) Get(id string) (domains.Item, error) {
	val, err := ms.md.Get(id)
	if err != nil {
		return domains.Item{}, fmt.Errorf("failed get item: %w", err)
	}
	return domains.Item{ID: id, Value: val}, nil
}

func (ms *memcachedStorage) Delete(id string) error {
	err := ms.md.Delete(id)
	if err != nil {
		return fmt.Errorf("failed delete item with id: %v, error: %w", id, err)
	}
	return nil
}
