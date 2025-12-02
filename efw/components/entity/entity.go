package entity

import (
	"sync"
	"time"
)

type BaseEntity[Id ~int64] struct {
	Id        Id    `json:"id" gorm:"column:id"`
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	mu sync.Mutex `json:"-" gorm:"-"`
}

func (b *BaseEntity[Id]) GetId() Id {
	return b.Id
}

func (b *BaseEntity[Id]) GetIdInt64() int64 {
	return int64(b.Id)
}

func (b *BaseEntity[Id]) SetId(id Id) {
	b.Id = id
}

func (b *BaseEntity[Id]) GetCreatedAt() int64 {
	return b.CreatedAt
}

func (b *BaseEntity[Id]) GetUpdatedAt() int64 {
	return b.UpdatedAt
}

func (b *BaseEntity[Id]) Lock() {
	b.mu.Lock()
}

func (b *BaseEntity[Id]) Unlock() {
	b.mu.Unlock()
}

func (b *BaseEntity[Id]) UpdateTimestamps() {
	now := time.Now().Unix()
	if b.CreatedAt == 0 {
		b.CreatedAt = now
	}
	b.UpdatedAt = now
}
