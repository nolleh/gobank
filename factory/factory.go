package factory

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"gobank/config"
	"gobank/echoMiddlewares"
	"gobank/logger"
	"google.golang.org/api/option"
)

func DB(ctx context.Context) xorm.Interface {
	v := ctx.Value(echoMiddlewares.ContextDBName)
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
	v := ctx.Value(logger.ContextLoggerName)
	if v == nil {
		return logrus.WithFields(logrus.Fields{})
	}
	if logger, ok := v.(*logrus.Entry); ok {
		return logger
	}
	return logrus.WithFields(logrus.Fields{})
}

func ApiContext(ctx context.Context) *echoMiddlewares.ApiContext {
	v := ctx.Value(echoMiddlewares.ApiContextName).(*echoMiddlewares.ApiContext)
	if v == nil {
		panic("not exist ApiContext")
	}
	return v
}

var DatastoreClient *datastore.Client

func NewDataStore(config *config.Config) *datastore.Client {
	if DatastoreClient == nil {
		ctx := context.Background()
		client, err := datastore.NewClient(ctx, config.DataStore.ProjectId,
			option.WithCredentialsFile(config.DataStore.KeyFile)); if err != nil {
			panic(err)
		}
		DatastoreClient = client
	}
	return DatastoreClient
}

func DataStore() *datastore.Client {
	if DatastoreClient == nil {
		panic("no datastore")
	}
	return DatastoreClient
}