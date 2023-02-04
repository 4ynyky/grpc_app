package inmemory_test

import (
	"strconv"
	"testing"

	"github.com/4ynyky/grpc_app/internal/domains"
	"github.com/4ynyky/grpc_app/internal/storage/inmemory"
)

func TestSetInMemory(t *testing.T) {
	tests := []struct {
		name     string
		args     interface{}
		expected string
		wantErr  bool
	}{
		{
			name:    "Positive",
			args:    domains.Item{ID: "1"},
			wantErr: false,
		},
		{
			name:    "Set item with empty ID",
			args:    domains.Item{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := inmemory.NewInMemoryStorage()
			if err := st.Set(tt.args.(domains.Item)); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkSetInMemory(b *testing.B) {
	st := inmemory.NewInMemoryStorage()
	for i := 0; i < b.N; i++ {
		if err := st.Set(domains.Item{ID: strconv.Itoa(i)}); err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkGetInMemory(b *testing.B) {
	st := inmemory.NewInMemoryStorage()
	for i := 0; i < b.N; i++ {
		if err := st.Set(domains.Item{ID: strconv.Itoa(i)}); err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := st.Get(strconv.Itoa(i)); err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}
