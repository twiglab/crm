package lua

import lua "github.com/yuin/gopher-lua"

func Trans2LValue(value any) lua.LValue {
	switch converted := value.(type) {
	case bool:
		return lua.LBool(converted)
	case int:
		return lua.LNumber(converted)
	case int32:
		return lua.LNumber(converted)
	case int64:
		return lua.LNumber(converted)
	case float32:
		return lua.LNumber(converted)
	case float64:
		return lua.LNumber(converted)
	case string:
		return lua.LString(converted)
	case []byte:
		return lua.LString(converted)
	case []string:
		var arr lua.LTable
		for _, item := range converted {
			arr.Append(lua.LString(item))
		}
		return &arr
	case []any:
		var arr lua.LTable
		for _, item := range converted {
			arr.Append(Trans2LValue(item))
		}
		return &arr
	case map[string]any:
		var table lua.LTable
		for key, item := range converted {
			table.RawSetH(lua.LString(key), Trans2LValue(item))
		}
		return &table
	case []map[string]any:
		var arr lua.LTable
		for _, item := range converted {
			arr.Append(Trans2LValue(item))
		}
		return &arr
	case nil:
		return lua.LNil
	}
	return lua.LString(value.(string))
}

func Trans2GoValue(value lua.LValue) any {
	switch val := value.(type) {
	case *lua.LNilType:
		return nil
	case lua.LBool:
		return bool(val)
	case lua.LNumber:
		return float64(val)
	case lua.LString:
		return val.String()
	case *lua.LTable:
		maxn := val.MaxN()
		// table
		if maxn == 0 {
			var table = make(map[string]any)
			val.ForEach(func(key, value lua.LValue) {
				table[key.String()] = Trans2GoValue(value)
			})
			return table
		}
		// array
		var arr = make([]any, 0, maxn)
		for i := 1; i <= maxn; i++ {
			arr = append(arr, Trans2GoValue(val.RawGetInt(i)))
		}
		return arr
	default:
		return val
	}
}
