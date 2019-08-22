package models

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"gobank/factory"
)
/** BalanceEntity2 ...
 * entity2 is store model to google/datastore
 */

const (
	Kind = "BANK/BALANCE"
	NameSpace ="dev"
)

type balanceDatastoreService struct {}

var BalanceDatastore balanceDatastoreService

func (ds *balanceDatastoreService) GetById(ctx context.Context, userId int64) (*BalanceEntity, error) {
	key := datastore.IDKey(Kind, userId, nil)
	key.Namespace = NameSpace
	var entity BalanceEntity
	if err := factory.DataStore().Get(ctx, key, &entity); err != nil {
		return &entity, err
	}
	return &entity, nil
}

func (ds *balanceDatastoreService) Create(ctx context.Context, entity2 *BalanceEntity) error {
	key := datastore.IDKey(Kind, entity2.UserId, nil)
	key.Namespace = NameSpace
	if _, err := factory.DataStore().Put(ctx, key, entity2); err != nil {
		return err
	}
	return nil
}

func (ds *balanceDatastoreService) Delete(ctx context.Context, userId int64) error {
	key := datastore.IDKey(Kind, userId, nil)
	key.Namespace = NameSpace
	if err := factory.DataStore().Delete(ctx, key); err != nil {
		return err
	}
	return nil
}

func (ds *balanceDatastoreService) Update(ctx context.Context, entity *BalanceEntity) error {
	key := datastore.IDKey(Kind, entity.UserId, nil)
	key.Namespace = NameSpace
	if _, err := factory.DataStore().Put(ctx, key, entity); err != nil {
		return err
	}
	return nil
}

func (ds *balanceDatastoreService) UpdateByRelatively(ctx context.Context, id int64, relVal *Balance, action BalanceAction) (*BalanceEntity, error) {
	res, err := ds.GetById(ctx, id); if err != nil {
		b := BalanceEntity{ UserId: id, Balance: *relVal }
		err := ds.Create(ctx, &b)
		return &b, err
	}
	newAmount := res.Balance.Amount + relVal.Amount
	if action == Withdraw {
		newAmount = res.Balance.Amount - relVal.Amount
	}
	strExpr := fmt.Sprint(newAmount, " ", relVal.Symbol)
	balance := Balance{ Amount: newAmount, Symbol: relVal.Symbol, Fraction: relVal.Fraction, StrExpr: strExpr }
	b := BalanceEntity{ UserId: id, Balance: balance }
	uperr := ds.Update(ctx, &b)
	return &b, uperr
}