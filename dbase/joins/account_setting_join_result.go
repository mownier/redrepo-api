package joins

import (
	"redrepo-api/dbase/tables"
	)

type AccountSettingJoinResult struct {
	tables.Account
	ConnectedToFacebook int `db:"connected_to_facebook"`
	ConnectedToTwitter 	int `db:"connected_to_twitter"`
	Verified 			int `db:"verified"`
}