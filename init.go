package muma

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/reusee/dscope"
	"github.com/reusee/starlarkutil"
	"go.starlark.net/starlark"
)

func init() {
	go func() {

		exePath, err := os.Executable()
		ce(err)
		exeDir := filepath.Dir(exePath)
		scriptPath := filepath.Join(exeDir, "muma.py")

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGUSR2)
		for {
			<-c

			script, err := os.ReadFile(scriptPath)
			if is(err, os.ErrNotExist) {
				continue
			}

			scope := dscope.New(dscope.Methods(new(Global))...)
			var fns ScriptFuncs
			scope.Assign(&fns)
			pyFuncs := make(starlark.StringDict)
			for name, fn := range fns {
				pyFuncs[name] = starlarkutil.MakeFunc(name, fn)
			}

			_, err = starlark.ExecFile(
				new(starlark.Thread),
				scriptPath,
				script,
				pyFuncs,
			)
			ce(err)

		}
	}()
}
