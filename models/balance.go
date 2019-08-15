/* Package models ...
* each model has it's definitions for how to store the data to storage.
*/
package models

import (
	"context"
	"gobank/factory"
	"time"
)

// Balance ...
type Balance struct {
	UserId int64 `json:"userId"`
	Amount int64 `json:"amount"`
	CreatedAt time.Time `json:"createdAt" xorm:"created`
	UpdatedAt time.Time `json:"updatedAt xorm:"updated`
}


// Create ... Insert To DB
func (b *Balance) Create(ctx context.Context) (int64, error) {
	return factory.DB(ctx).Insert(b)
	// return 0, nil
}

// Delete ... From DB
func (b *Balance) Delete(ctx context.Context) (int64, error) {
	return factory.DB(ctx).Delete(b)
}

// Update ... From DB
func (b *Balance) Update(ctx context.Context) (int64, error) {
	return factory.DB(ctx).ID(b.UserId).Update(b)
}