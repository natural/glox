package glox

import (
	"io"
	"net/http"

	"github.com/yuin/gopher-lua"
)

//
//
func LHttpRequest(L *lua.LState, r *http.Request) *lua.LTable {
	t := L.NewTable()
	t.RawSetString("body", LReadCloser(L, r.Body))
	t.RawSetString("close", lua.LBool(r.Close))
	t.RawSetString("content_length", lua.LNumber(r.ContentLength))
	t.RawSetString("header", LHttpHeader(L, r.Header))
	t.RawSetString("host", lua.LString(r.Host))
	t.RawSetString("method", lua.LString(r.Method))
	t.RawSetString("proto", lua.LString(r.Proto))
	t.RawSetString("proto_major", lua.LNumber(r.ProtoMajor))
	t.RawSetString("proto_minor", lua.LNumber(r.ProtoMinor))
	t.RawSetString("remote_addr", lua.LString(r.RemoteAddr))
	t.RawSetString("request_uri", lua.LString(r.RequestURI))
	t.RawSetString("trailer", LHttpHeader(L, r.Trailer))
	t.RawSetString("tls", lua.LNil)
	// Cancel <-chan struct{} // not applicable to server per net/http docs

	// how to make lazy?  func or...?
	t.RawSetString("transfer_encoding", LStringSlice(L, r.TransferEncoding))

	// this could be much better
	t.RawSetString("url", lua.LString(r.URL.String()))

	L.SetFuncs(t, map[string]lua.LGFunction{
		// AddCookie(c *Cookie)
		// BasicAuth() (username, password string, ok bool)
		// Cookie(name string) (*Cookie, error)
		// Cookies() []*Cookie
		// FormFile(key string) (multipart.File, *multipart.FileHeader, error)
		// FormValue(key string) string
		// MultipartReader() (*multipart.Reader, error)
		// ParseForm() error
		// ParseMultipartForm(maxMemory int64) error
		// PostFormValue(key string) string
		// ProtoAtLeast(major, minor int) bool
		// Referer() string
		// SetBasicAuth(username, password string)
		"user_agent": func(L *lua.LState) int {
			L.Push(lua.LString(r.UserAgent()))
			return 1
		},
		// Write(w io.Writer) error
		"write": func(L *lua.LState) int {
			o := L.CheckAny(1)
			w, ok := o.(io.Writer)
			if !ok {
				L.Push(lua.LString("not a writer"))
				return 1
			}
			err := r.Write(w)
			if err == nil {
				L.Push(lua.LNil)
			} else {
				L.Push(lua.LString(err.Error()))
			}
			return 1
		},
		// WriteProxy(w io.Writer) error
		"write_proxy": func(L *lua.LState) int {
			o := L.CheckAny(1)
			w, ok := o.(io.Writer)
			if !ok {
				L.Push(lua.LString("not a writer"))
				return 1
			}
			err := r.WriteProxy(w)
			if err == nil {
				L.Push(lua.LNil)
			} else {
				L.Push(lua.LString(err.Error()))
			}
			return 1
		},
	})
	return t
}

// make the "response" api (incomplete)
//
func LHttpResponseWriter(L *lua.LState, w http.ResponseWriter) *lua.LTable {
	t := L.NewTable()
	L.SetFuncs(t, map[string]lua.LGFunction{

		// Write([]byte) (int, error)
		// write(string) (int, string or nil)
		"write": func(L *lua.LState) int {
			s := L.CheckString(1)
			if n, err := w.Write([]byte(s)); err == nil {
				L.Push(lua.LNumber(n))
				L.Push(lua.LNil)
			} else {
				L.Push(lua.LNumber(0))
				L.Push(lua.LString(err.Error()))
			}
			return 2
		},

		// WriteHeader(int)
		// write_header(int)
		"write_header": func(L *lua.LState) int {
			w.WriteHeader(L.CheckInt(1))
			return 0
		},

		// Header() Header
		//
		"header": func(L *lua.LState) int {
			L.Push(LHttpHeader(L, w.Header()))
			return 1
		},
	})
	return t
}

//
//
//type Header map[string][]string
func LHttpHeader(L *lua.LState, h http.Header) *lua.LTable {
	t := L.NewTable()
	//	for k, vs := range h {
	//		t.RawSetString(k, LStringSlice(L, vs))
	//	}
	L.SetFuncs(t, map[string]lua.LGFunction{
		//func (h Header) Add(key, value string)
		"add": func(L *lua.LState) int {
			h.Add(L.CheckString(1), L.CheckString(2))
			return 0
		},
		//func (h Header) Del(key string)
		"del": func(L *lua.LState) int {
			h.Del(L.CheckString(1))
			return 0
		},
		//func (h Header) Get(key string) string
		"get": func(L *lua.LState) int {
			L.Push(lua.LString(h.Get(L.CheckString(1))))
			return 1
		},
		//func (h Header) Set(key, value string)
		"set": func(L *lua.LState) int {
			h.Set(L.CheckString(1), L.CheckString(2))
			return 0
		},
		//func (h Header) Write(w io.Writer) error
		"write": func(L *lua.LState) int {
			L.Push(lua.LNil)
			return 1
		},
		//func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
		"write_subset": func(L *lua.LState) int {
			L.Push(lua.LNil)
			return 1
		},
	})
	return t
}

//
//
func LOpenGoGlobals(L *lua.LState) error {
	//error makes sense as a type instead
	//append doesn't make sense
	//delete doesn't make sense
	//length  doesn't make sense
	return nil
}
