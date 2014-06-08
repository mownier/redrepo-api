// config.go
// @author Mounir Ybanez
// @date June 8, 2014

package dbase

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "github.com/coopernurse/gorp"
)

func OpenDatabase() *gorp.DbMap {
	// connect to db using standard Go database/sql API
    // use whatever database/sql driver you wish
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/redrepo")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	return dbmap
}

func CloseDatabase(dbmap *gorp.DbMap) {
	dbmap.Db.Close()
}