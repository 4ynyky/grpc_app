package services

import (
	"github.com/4ynyky/grpc_app/pkg/domains"
	"github.com/4ynyky/grpc_app/pkg/storage"
	"github.com/sirupsen/logrus"
)

type IStorageService interface {
	Set(item domains.Item) error
	Get(id string) (domains.Item, error)
	Delete(id string) error
}

type storageService struct {
	storage storage.IStorage
}

func NewStorageService(storage storage.IStorage) IStorageService {
	return &storageService{storage: storage}
}

func (ss *storageService) Set(item domains.Item) error {
	if err := ss.storage.Set(item); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
func (ss *storageService) Get(id string) (domains.Item, error) {
	item, err := ss.storage.Get(id)
	if err == storage.ErrNotFound {
		logrus.Warnf("Failed get item with id %v, error: %v", id, err)
		return domains.Item{}, err
	} else if err != nil {
		logrus.Errorf("Failed get item with id %v, error: %w", id, err)
		return domains.Item{}, err
	}
	return item, nil
}
func (ss *storageService) Delete(id string) error {
	if err := ss.storage.Delete(id); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
