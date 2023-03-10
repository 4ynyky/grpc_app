package inmemory

import (
	"sync"

	"github.com/4ynyky/grpc_app/internal/domains"
	"github.com/4ynyky/grpc_app/internal/storage"
)

type inmemoryStorage struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewInMemoryStorage() *inmemoryStorage {
	return &inmemoryStorage{data: make(map[string]string)}
}

func (is *inmemoryStorage) Set(item domains.Item) error {
	if len(item.ID) == 0 {
		return storage.ErrItemInvalid
	}
	is.mu.Lock()
	defer is.mu.Unlock()
	is.data[item.ID] = item.Value
	return nil
}

func (is *inmemoryStorage) Get(id string) (domains.Item, error) {
	is.mu.RLock()
	defer is.mu.RUnlock()
	val, ok := is.data[id]
	if !ok {
		return domains.Item{}, storage.ErrNotFound
	}
	return domains.Item{ID: id, Value: val}, nil
}

func (is *inmemoryStorage) Delete(id string) error {
	is.mu.Lock()
	defer is.mu.Unlock()
	delete(is.data, id)
	return nil
}
