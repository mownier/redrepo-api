// redrepo-api.go
// @author Mounir Ybanez
// @date June 7, 2014

package main

import (
	api "redrepo-api/services"
)

func init() {
	api.Start(2121)
}