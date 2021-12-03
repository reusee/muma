package muma

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func (_ Global) Profiles() ScriptFuncs {
	var f *os.File
	var err error

	profileFunc := func(name string) func() {
		return func() {
			profile := pprof.Lookup(name)
			if profile == nil {
				return
			}
			f, err := os.Create(fmt.Sprintf("%s-profile-%v", name, time.Now()))
			ce(err)
			defer f.Close()
			ce(profile.WriteTo(f, 0))
		}
	}

	return ScriptFuncs{

		"cpu_profile": func() {
			if f == nil {
				// start
				f, err = os.Create(fmt.Sprintf("cpu-profile-%v", time.Now()))
				ce(err)
				pprof.StartCPUProfile(f)
			} else {
				// stop
				pprof.StopCPUProfile()
				ce(f.Close())
				f = nil
			}
		},

		"goroutine_profile":    profileFunc("goroutine"),
		"heap_profile":         profileFunc("heap"),
		"allocs_profile":       profileFunc("allocs"),
		"threadcreate_profile": profileFunc("threadcreate"),
		"block_profile":        profileFunc("block"),
		"mutex_profile":        profileFunc("mutex"),
	}
}
