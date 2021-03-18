package database

import (
	"database/sql"
	"log"
	"time"
)

// DbConn の DB type は コネクションのプール。スレッドセーフ
// Export できるよう、大文字で
var DbConn *sql.DB

// SetupDatabase は。。。
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:jA3p1onn@tcp(127.0.0.1:3306)/characterdb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(3)
	// コネクションをプールしたとき、アイドルなコネクションがいくつ残っていることを許容するか
	DbConn.SetMaxIdleConns(3)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
