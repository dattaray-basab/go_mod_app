package main

import (
	"github.com/dattaray-basab/go_mod"
	"github.com/dattaray-basab/go_mod/pkg1"
	"github.com/dattaray-basab/go_mod/pkg2"
)

func main() {
	go_mod.Hello()

	pkg1.F1()
	pkg2.F2()
	pkg2.F2_1()
}
