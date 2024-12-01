package lua

import (
	lua "github.com/yuin/gopher-lua"
	"time"
)

const (
	oneDayMSec = 86400 * 1000
)

func LuaModuleLoader(L *lua.LState) int {
	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}

var api = map[string]lua.LGFunction{
	"getDayBeforeTimeStamp": getDayBeforeTimeStamp,
}

func getDayBeforeTimeStamp(L *lua.LState) int {
	curSec := time.Now().UnixMilli()
	day := L.CheckNumber(1)
	zeroClockMSec := curSec - curSec%oneDayMSec
	dayBeforeMSec := zeroClockMSec - int64(day)*oneDayMSec
	L.Push(Trans2LValue(dayBeforeMSec))
	return 1
}
