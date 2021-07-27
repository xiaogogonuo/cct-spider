package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/xiaogogonuo/cct-spider/pkg/config"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"strconv"
	"time"
)

const (
	queryTimeOut time.Duration = 30
	transactionTimeOut time.Duration = 60
)

var db *sql.DB

func dbConfig() *viper.Viper {
	c := config.Config{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: "configs/db",
	}
	v, err := c.NewConfig()
	if err != nil {
		logger.DPanic(err.Error())
	}
	return v
}

func init() {
	v := dbConfig()
	mysql := v.GetStringMapString("mysql")

	// 用户名:密码@协议(ip:端口)/数据库?charset=utf8&parseTime=true&loc=Local
	dataSourceName := mysql["user"] + ":" + mysql["pass"] + "@tcp(" + mysql["host"] + ":" + mysql["port"] + ")/" +
		mysql["db"] + "?charset=utf8&loc=Local"

	_db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.DPanic(err.Error())
	}

	maxOpenConn, err := strconv.Atoi(mysql["maxopenconn"])
	if err != nil {
		logger.DPanic(err.Error())
	}
	_db.SetMaxOpenConns(maxOpenConn)

	maxIdleConn, err := strconv.Atoi(mysql["maxidleconn"])
	if err != nil {
		logger.DPanic(err.Error())
	}
	_db.SetMaxIdleConns(maxIdleConn)

	if err = _db.Ping(); err != nil {
		logger.DPanic(err.Error())
	}

	db = _db
}

func scan(query string, result chan []string, stop chan struct{}) {
	rows, err := db.Query(query)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	vars := make([]sql.NullString, len(cols))
	dest := make([]interface{}, len(cols))
	for idx := range vars {
		dest[idx] = &vars[idx]
	}
	for rows.Next() {
		// 结果集方法Scan可以把数据库取出的字段值赋值给指定的数据结构
		if err = rows.Scan(dest...); err != nil {
			continue
		}
		var m []string
		for _, v := range vars {
			if v.Valid {
				m = append(m, v.String)
			} else {
				m = append(m, "")
			}
		}
		result <- m
	}
	stop <- struct{}{}
	close(stop)
	close(result)
	return
}

func query(ctx context.Context, sql string, result chan []string) {
	stop := make(chan struct{})
	go scan(sql, result, stop)
	select {
	case <-ctx.Done():
	case <-stop:
	}
}

// Query return rows within queryTimeOut
func Query(sql string) (row [][]string) {
	result := make(chan []string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*queryTimeOut)
	defer cancel()
	go query(ctx, sql, result)
	for {
		select {
		case <-ctx.Done():
			logger.Error(ctx.Err().Error())
			return
		case res, ok := <-result:
			if !ok {
				return
			}
			row = append(row, res)
		}
	}
}

func clearTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		logger.Error(err.Error())
	}
}

func exec(sql string, stop chan struct{}, data ...interface{}) {
	defer close(stop)
	tx, err := db.Begin()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer clearTransaction(tx)
	r, err := tx.Exec(sql, data...)
	if err != nil {
		logger.Error(err.Error() + `\n` + sql, logger.Field("data", data))
		return
	}
	if _, err = r.RowsAffected(); err != nil {
		logger.Error(err.Error())
		return
	}
	if err = tx.Commit(); err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("Insert success")
	return
}

func transaction(ctx context.Context, sql string, sig chan struct{}, data ...interface{}) {
	stop := make(chan struct{})
	go exec(sql, stop, data...)
	select {
	case <-ctx.Done():
	case <-stop:
		sig<- struct{}{}
	}
}

func Transaction(sql string, data ...interface{}) {
	sig := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*transactionTimeOut)
	defer cancel()
	go transaction(ctx, sql, sig, data...)
	select {
	case <-ctx.Done():
	case <-sig:
	}
}