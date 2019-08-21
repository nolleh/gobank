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
	Amount int64 `json:"amount"`
	Symbol string `json:"symbol"`
	Fraction int8 `json:"fraction"`
	StrExpr string `json:"strExpr"`
}

// Balance ...
type BalanceEntity struct {
	UserId int64 `json:"userId" xorm:"pk"`
	Balance Balance `json:"balance" xorm:"balance json notnull"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}

type BalanceRDB struct {
}

var BalanceMySql BalanceRDB

func (dbService *BalanceRDB) GetById(ctx context.Context, userId int64) (*BalanceEntity, error) {
	 var b BalanceEntity
	 _, err := factory.DB(ctx).ID(userId).Get(b)
	 return &b, err
}

// Create ... Insert To DB
func (dbService *BalanceRDB) Create(ctx context.Context, entity *BalanceEntity) (*BalanceEntity, error) {
	var b BalanceEntity
	_, err := factory.DB(ctx).Insert(b)
	return &b, err
}

// Delete ... From DB
func (dbService *BalanceRDB) Delete(ctx context.Context, userId int64) (*BalanceEntity, error) {
	var b BalanceEntity
	_, err := factory.DB(ctx).ID(userId).Delete(b)
	return &b, err
}

// Update ... From DB
func (dbService *BalanceRDB) Update(ctx context.Context, entity *BalanceEntity) (*BalanceEntity, error) {
	_, err := factory.DB(ctx).ID(entity.UserId).Update(entity)
	return entity, err
}

func (dbService *BalanceRDB) UpdateByRelatively(ctx context.Context, userId int64, relVal Balance) (*BalanceEntity, error) {
	res, err := dbService.GetById(ctx, userId); if err != nil {
		b := &BalanceEntity{ UserId: userId, Balance: relVal}
		return dbService.Create(ctx, b)
	}
	newAmount := res.Balance.Amount + relVal.Amount
	strExpr := fmt.Sprint(newAmount, " ", res.Balance.Symbol)
	balance := Balance {Amount: newAmount, Symbol: res.Balance.Symbol, Fraction: res.Balance.Fraction, StrExpr: strExpr}
	res.Balance = balance
	_, err = factory.DB(ctx).ID(res.UserId).Update(res)
	return res, err
}