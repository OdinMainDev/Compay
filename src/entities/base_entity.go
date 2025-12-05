package entities

import (
	"sync"
	"time"
)

// BaseEntity provides common fields and concurrency primitives for models.
type BaseEntity[ID ~int64] struct {
	ID        ID    `json:"id" gorm:"column:id"`
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	mu sync.Mutex `json:"-" gorm:"-"`
}

func (b *BaseEntity[ID]) TypedID() ID {
	return b.ID
}

func (b *BaseEntity[ID]) GetId() int64 {
	return int64(b.ID)
}

func (b *BaseEntity[ID]) SetId(id ID) {
	b.ID = id
}

func (b *BaseEntity[ID]) GetCreatedAt() int64 {
	return b.CreatedAt
}

func (b *BaseEntity[ID]) GetUpdatedAt() int64 {
	return b.UpdatedAt
}

func (b *BaseEntity[ID]) Lock() {
	b.mu.Lock()
}

func (b *BaseEntity[ID]) Unlock() {
	b.mu.Unlock()
}

// TouchTimestamps updates CreatedAt/UpdatedAt prior to persisting a record.
func (b *BaseEntity[ID]) TouchTimestamps() {
	now := time.Now().Unix()
	if b.CreatedAt == 0 {
		b.CreatedAt = now
	}
	b.UpdatedAt = now
}
