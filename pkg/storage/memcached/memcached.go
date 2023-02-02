package memcached

import (
	"fmt"

	"github.com/4ynyky/grpc_app/pkg/domains"
	"github.com/4ynyky/grpc_app/pkg/storage"
	"github.com/bradfitz/gomemcache/memcache"
)

type Config struct {
	Host string
}

type memcachedStorage struct {
	client *memcache.Client
}

func NewMemcachedStorage(conf Config) (storage.IStorage, error) {
	ms := &memcachedStorage{}
	ms.client = memcache.New(conf.Host)

	if ms.client == nil {
		return nil, storage.ErrConnNotEstablished
	}
	return ms, nil
}

func (ms *memcachedStorage) Set(item domains.Item) error {
	if err := ms.client.Set(&memcache.Item{Key: item.ID, Value: []byte(item.Value)}); err != nil {
		return fmt.Errorf("failed store item: %v, error: %w", item, err)
	}
	return nil
}

func (ms *memcachedStorage) Get(id string) (domains.Item, error) {
	memItem, err := ms.client.Get(id)
	if err == memcache.ErrCacheMiss {
		return domains.Item{}, storage.ErrNotFound
	} else if err != nil {
		return domains.Item{}, fmt.Errorf("failed get item with id: %v, error: %w", id, err)
	}
	return domains.Item{ID: memItem.Key, Value: string(memItem.Value)}, nil
}

func (ms *memcachedStorage) Delete(id string) error {
	err := ms.client.Delete(id)
	if err != nil {
		return fmt.Errorf("failed delete item with id: %v, error: %w", id, err)
	}
	return nil
}
