package glox

import (
	"io"
	"io/ioutil"

	"github.com/yuin/gopher-lua"
)

//
//
func LReadCloser(L *lua.LState, o io.ReadCloser) *lua.LTable {
	t := L.NewTable()
	L.SetFuncs(t, map[string]lua.LGFunction{
		// Read(p []byte) (n int, err error)
		// read() (string, string or nil)
		"read": func(L *lua.LState) int {
			if bs, err := ioutil.ReadAll(o); err == nil {
				L.Push(lua.LString(string(bs)))
				L.Push(lua.LNil)

			} else {
				L.Push(lua.LString(""))
				L.Push(lua.LString(err.Error()))
			}
			return 2
		},

		// Close() error
		// close() (string or nil)
		"close": func(L *lua.LState) int {
			if err := o.Close(); err == nil {
				L.Push(lua.LNil)
			} else {
				L.Push(lua.LString(err.Error()))
			}
			return 1
		},
	})
	return t
}
