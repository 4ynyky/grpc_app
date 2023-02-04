package memcached

import (
	"strconv"
	"testing"

	"github.com/4ynyky/grpc_app/internal/domains"
)

func BenchmarkSetMemdriver(b *testing.B) {
	ms, err := NewMemcachedStorage(Config{Host: "0.0.0.0:11211"})
	if err != nil {
		b.Fatalf("For start this benchmark for first start memcache: %v", err)
	}
	for i := 0; i < b.N; i++ {
		if err := ms.Set(domains.Item{ID: strconv.Itoa(i)}); err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}
