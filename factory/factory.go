package factory

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"github.com/nolleh/gobank/logger"
)

func DB(ctx context.Context) xorm.Interface {
	v := ctx.Value(ContextDBName)
	if v == nil {
		panic("DB is not exist")
	}

	if db, ok := v.(*xorm.Session); ok {
		return db
	}

	if db, ok := v.(*xorm.Engine); ok {
		return db
	}
	panic("DB is not exist")
}

func Logger(ctx context.Context) *logrus.Entry {
	v := ctx.Value(loggers.ContextLoggerName)
	if v == nil {
		return logrus.WithFields(logrus.Fields{})
	}
	if logger, ok := v.(*logrus.Entry); ok {
		return logger
	}
	return logrus.WithFields(logrus.Fields{})
}
