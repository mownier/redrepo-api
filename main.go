// api.go
// @author Mounir Ybanez
// @date May 6, 2014

package main

import (
	api "redrepo-api/services"
)

func main() {
	api.Start(2121)
}
