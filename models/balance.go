package models

import (
	"context"
	"time"
)

// balance model
type Balance struct {
	Id int64 `json:"id"`
	CreatedAt time.Time `json:"createdAt" xorm:"created`
	UpdatedAt time.Time `json:"updatedAt xorm:"updated`
}

func (b *Balance) Create(ctx context.Context) (int64, error) {
	return factory.DB(ctx).Insert(b)
}