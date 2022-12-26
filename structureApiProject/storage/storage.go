package storage

import (
	"GoExpGain/structureApiProject/types"
)

type Storage interface {
	Get(int) *types.User
}