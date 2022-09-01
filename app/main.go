package main

import (
	"github.com/jhonatanlteodoro/verify_test/app/initalization"
)

func main() {
	a := initalization.App{}
	a.Initilize()
	a.Run("", "8080")
}
