package storage

import "github.com/4ynyky/grpc_app/internal/domains"

type IStorage interface {
	Set(item domains.Item) error
	Get(id string) (domains.Item, error)
	Delete(id string) error
}
