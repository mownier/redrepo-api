// config.go
// @author Mounir Ybanez
// @date June 8, 2014

package dbase

import (
	"database/sql"
	"redrepo-api/dbase/tables"
	_ "github.com/go-sql-driver/mysql"
    "github.com/coopernurse/gorp"
    "fmt"
    "errors"
)

func OpenDatabase() (*gorp.DbMap, error) {
	// connect to db using standard Go database/sql API
    // use whatever database/sql driver you wish
    db, err := sql.Open("mysql", "mownier:mownier@tcp(localhost:3306)/redrepo")
	if err != nil {
		return nil, errors.New("Cannot establish database connection.")
	} else {
		fmt.Println("Established database connection.")
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		return nil, errors.New("DSN data validation failed.")
	} else {
		fmt.Println("DSN data validation success.")
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(tables.Account{}, "accounts")
	dbmap.AddTableWithName(tables.AccountSetting{}, "account_settings")
	dbmap.AddTableWithName(tables.VerificationCode{}, "verification_codes")
	dbmap.AddTableWithName(tables.AuthClient{}, "auth_clients")
	dbmap.AddTableWithName(tables.AuthSession{}, "auth_sessions")
	
	return dbmap, nil
}

func CloseDatabase(dbmap *gorp.DbMap) {
	dbmap.Db.Close()
}