package main

import (
	"fmt"
	"github.com/twiglab/crm/poly/lua"
)

type Member struct {
}

func (ins *Member) GetMemberLevel() int {
	return 1
}

func main() {
	scriptPath := "script/test.lua"

	runner := lua.NewLuaRunner()
	luaProto, err := runner.CompileLuaFile(scriptPath)
	if err != nil {
		return
	}
	compL, err := runner.NewLuaState(luaProto)
	if err != nil {
		return
	}
	// 注册go模块
	runner.RegisterGoModule(map[string]any{
		"member": new(Member),
	})
	// input参数
	inputParam := map[string]any{
		"id":           1,
		"client_level": 1,
	}
	resultMap, err := runner.RunLua(compL, lua.Trans2LValue(inputParam))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultMap)
}
