package social_graph_manager

import (
	"errors"

	om "github.com/yuan8180/delinkcious/pkg/object_model"
)

type MilvusManager struct {
	store om.MilvusManager
}

func NewMilvusManager(store om.MilvusManager) (om.MilvusManager, error) {
	if store == nil {
		return nil, errors.New("store can't be nil")
	}
	return &MilvusManager{store: store}, nil
}

func (m *MilvusManager) CreateCollection(collname string) (err error) {
	if collname == "" {
		err = errors.New("collname can't be empty")
		return
	}

	return m.store.CreateCollection(collname)
}

func (m *MilvusManager) DropCollection(collname string) (err error) {
	if collname == "" {
		err = errors.New("collname can't be empty")
		return
	}

	return m.store.DropCollection(collname)
}

func (m *MilvusManager) ListCollections() ([]string, error) {
	return m.store.ListCollections()
}

func (m *MilvusManager) HasCollection(collname string) (collExists bool, err error) {
	if collname == "" {
		err = errors.New("collname can't be empty")
		return
	}
	collExists, err = m.store.HasCollection(collname)
	return
}
