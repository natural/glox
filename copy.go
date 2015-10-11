package glox

import (
	"github.com/yuin/gopher-lua"
)

//
//
func LCopyGlobal(src *lua.LState, dst *lua.LState, keys ...string) {
	for _, k := range keys {
		dst.SetGlobal(k, src.GetGlobal(k))
	}
}
