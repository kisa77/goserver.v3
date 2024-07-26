package main

import (
	_ "github.com/kisa77/goserver.v3/mmo"

	"github.com/kisa77/goserver.v3/core"
	"github.com/kisa77/goserver.v3/core/module"
)

func main() {
	defer core.ClosePackages()
	core.LoadPackages("config.json")

	waiter := module.Start()
	waiter.Wait("main")
}
