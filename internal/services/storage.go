package services

import (
	"github.com/4ynyky/grpc_app/pkg/domains"
	"github.com/4ynyky/grpc_app/pkg/storage"
)

type IStorageService interface {
	Set(id string, item domains.Item) error
	Get(id string) (domains.Item, error)
	Delete(id string) error
}

type storageService struct {
	storage storage.IStorage
}

func NewStorageService(storage storage.IStorage) IStorageService {
	return &storageService{storage: storage}
}

func (ss *storageService) Set(id string, item domains.Item) error {
	return nil
}
func (ss *storageService) Get(id string) (domains.Item, error) {
	return domains.Item{}, nil
}
func (ss *storageService) Delete(id string) error {
	return nil
}
