/* Package models ...
* each model has it's definitions for how to store the data to storage.
*/
package models

import (
	"context"
	"fmt"
	"gobank/factory"
	"time"
)

type Balance struct {
	Amount uint64 `json:"amount"`
	Symbol string `json:"symbol"`
	Fraction uint8 `json:"fraction"`
	StrExpr string `json:"strExpr"`
}

// Balance ...
type BalanceEntity struct {
	UserId uint64 `json:"userId" xorm:"pk"`
	Balance Balance `json:"balance" xorm:"balance json notnull"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}

func (b *BalanceEntity) GetById(ctx context.Context, id uint64) (bool, error) {
	return factory.DB(ctx).ID(id).Get(b)
}

// Create ... Insert To DB
func (b *BalanceEntity) Create(ctx context.Context) (int64, error) {
	return factory.DB(ctx).Insert(b)
	// return 0, nil
}

// Delete ... From DB
func (b *BalanceEntity) Delete(ctx context.Context) (int64, error) {
	return factory.DB(ctx).ID(b.UserId).Delete(b)
}

// Update ... From DB
func (b *BalanceEntity) Update(ctx context.Context) (int64, error) {
	return factory.DB(ctx).ID(b.UserId).Update(b)
}

func (b *BalanceEntity) UpdateByRelatively(ctx context.Context, relVal Balance) (int64, error) {
	if res, err := b.GetById(ctx, b.UserId); !res || err != nil {
		b = &BalanceEntity{ UserId: b.UserId, Balance: relVal}
		return b.Create(ctx)
	}
	newAmount := b.Balance.Amount + relVal.Amount
	strExpr := fmt.Sprint(newAmount, " ", b.Balance.Symbol)
	balance := Balance {Amount: newAmount, Symbol: b.Balance.Symbol, Fraction: b.Balance.Fraction, StrExpr: strExpr}
	b.Balance = balance
	return factory.DB(ctx).ID(b.UserId).Update(b)
}