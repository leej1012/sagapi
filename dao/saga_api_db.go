package dao

import (
	"context"
	"database/sql"
	"github.com/ontio/sagapi/config"
	"time"
)

var DefSagaApiDB *SagaApiDB

type SagaApiDB struct {
	ApiDB   *ApiDB
	OrderDB *OrderDB
}

func NewSagaApiDB(sagaDBConfig *config.Config) (*SagaApiDB, error) {
	db, dberr := sql.Open("mysql",
		sagaDBConfig.DbConfig.ProjectDBUser+
			":"+sagaDBConfig.DbConfig.ProjectDBPassword+
			"@tcp("+sagaDBConfig.DbConfig.ProjectDBUrl+
			")/"+sagaDBConfig.DbConfig.ProjectDBName+
			"?charset=utf8")
	if dberr != nil {
		return nil, dberr
	}
	ctx, cf := context.WithTimeout(context.Background(), 10*time.Second)
	defer cf()
	err := db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return &SagaApiDB{
		ApiDB:   NewApiDB(db),
		OrderDB: NewOrderDB(db),
	}, nil
}

func (this *SagaApiDB) Close() error {
	return this.ApiDB.db.Close()
}
