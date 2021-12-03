package muma

import (
	"reflect"

	"github.com/reusee/dscope"
)

type ScriptFuncs map[string]any

var _ dscope.Reducer = ScriptFuncs{}

func (_ ScriptFuncs) Reduce(_ dscope.Scope, vs []reflect.Value) reflect.Value {
	return dscope.Reduce(vs)
}

func (_ Global) ScriptFuncs() ScriptFuncs {
	return nil
}
