// api.go
// @author Mounir Ybanez
// @date May 6, 2014

package main

import (
	redrepoAPI "redrepo-api/services"
)

func main() {
	redrepoAPI.Start(2121)
}
