package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ontio/sagapi/sagaconfig"
)

type ApiKeyDB struct {
	db *sql.DB
}

func NewApiKeyDB(db *sql.DB) *ApiKeyDB {
	return &ApiKeyDB{
		db: db,
	}
}

func (this *ApiKeyDB) UpdateOrderStatusInApiKey(orderId string, status sagaconfig.OrderStatus) error {
	strSql := "update tbl_api_key set OrderStatus=? where OrderId=?"
	stmt, err := this.db.Prepare(strSql)
	if stmt != nil {
		defer stmt.Close()
	}
	if err != nil {
		return err
	}
	_, err = stmt.Exec(status, orderId)
	return err
}
