package storage

import "github.com/4ynyky/grpc_app/pkg/domains"

type IStorage interface {
	Set(item domains.Item) error
	Get(id string) (domains.Item, error)
	Delete(id string) error
}
