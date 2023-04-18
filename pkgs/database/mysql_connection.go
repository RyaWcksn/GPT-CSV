package database

import (
	"database/sql"
	"time"

	"github.com/RyaWcksn/nann-e/config"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	MYSQL config.Config
	L     logger.ILogger
}

func NewDatabaseConnection(M config.Config, l logger.ILogger) *Connection {
	return &Connection{
		MYSQL: M,
		L:     l,
	}
}

func (db *Connection) DBConnect() *sql.DB {
	dbConn, errConn := sql.Open(
		"mysql", db.MYSQL.Database.Username+":"+db.MYSQL.Database.Password+"@tcp("+db.MYSQL.Database.Host+")/"+db.MYSQL.Database.Database,
	)
	if errConn != nil {
		db.L.Errorf("[ERR] Error while connecting... := %v", errConn)
		return nil
	}
	for dbConn.Ping() != nil {
		db.L.Info("Attempting connect to DB...")
		time.Sleep(5 * time.Second)
	}
	dbConn.SetMaxIdleConns(db.MYSQL.Database.MaxIdleConn)
	dbConn.SetMaxOpenConns(db.MYSQL.Database.MaxOpenConn)
	return dbConn
}
