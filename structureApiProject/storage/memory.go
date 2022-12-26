package storage

import "GoExpGain/structureApiProject/types"

type MemoryStorage struct {}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *types.User {
	return &types.User{
		ID: 2,
		Name: "Brr",
	}
}