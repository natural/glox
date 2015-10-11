package glox

import (
	"github.com/yuin/gopher-lua"
)

//
//
func LStringSlice(L *lua.LState, vs []string) *lua.LTable {
	t := L.NewTable()
	for _, v := range vs {
		t.Append(lua.LString(v))
	}
	return t
}
