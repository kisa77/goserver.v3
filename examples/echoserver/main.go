// main
package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/kisa77/goserver.v3/core"
	_ "github.com/kisa77/goserver.v3/core/builtin/filter"
	"github.com/kisa77/goserver.v3/core/module"
)

func main() {
	defer core.ClosePackages()
	core.LoadPackages("config.json")
	//usage: go tool pprof http://localhost:6060/debug/pprof/heap
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	waiter := module.Start()
	waiter.Wait("main")
}
