package lua

import (
	"bufio"
	"errors"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	lua_parse "github.com/yuin/gopher-lua/parse"
	luar "layeh.com/gopher-luar"
	"os"
	"strings"
)

type LuaRunner struct {
	moduleMap map[string]any
}

func NewLuaRunner() *LuaRunner {
	ins := &LuaRunner{}
	ins.initRunner()
	return ins
}

func (ins *LuaRunner) initRunner() {
	ins.moduleMap = make(map[string]any)
}

func (ins *LuaRunner) RegisterGoModule(moduleMap map[string]any) {
	ins.moduleMap = moduleMap
}

func (ins *LuaRunner) RunLua(compL *lua.LState, inputTable lua.LValue) (any, error) {
	defer compL.Close()
	//
	ins.execRegModule(compL)
	if err := compL.CallByParam(lua.P{
		Fn:      compL.GetGlobal("run"),
		NRet:    lua.MultRet,
		Protect: true,
	}, inputTable); err != nil {
		return nil, fmt.Errorf("LuaRunner CallByParam err:%w", err)
	}
	dataRet := compL.Get(1)
	if dataRet == lua.LNil {
		return nil, errors.New("LuaRunner dataResult is nil")
	}
	resultRet := compL.Get(2)
	if resultRet == lua.LNil {
		return nil, errors.New("LuaRun errorResult is nil")
	}
	compL.Pop(2)
	if len(resultRet.String()) > 0 {
		return nil, errors.New(resultRet.String())
	}
	return Trans2GoValue(dataRet), nil
}

func (ins *LuaRunner) execRegModule(comp *lua.LState) {
	for name, module := range ins.moduleMap {
		comp.SetGlobal(name, luar.New(comp, module))
	}
}

func (ins *LuaRunner) CompileLuaString(codestr string) (*lua.FunctionProto, error) {
	reader := strings.NewReader(codestr)
	chunk, err := lua_parse.Parse(reader, codestr)
	if err != nil {
		return nil, fmt.Errorf("lua_parse.Parse err:%w", err)
	}
	proto, err := lua.Compile(chunk, codestr)
	if err != nil {
		return nil, fmt.Errorf("lua.Compile err:%w", err)
	}
	return proto, nil
}

func (ins *LuaRunner) CompileLuaFile(filePath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := lua_parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

func (ins *LuaRunner) doCompileFunc(L *lua.LState, proto *lua.FunctionProto) error {
	lfunc := L.NewFunctionFromProto(proto)
	L.Push(lfunc)
	return L.PCall(0, lua.MultRet, nil)
}

func (ins *LuaRunner) NewLuaState(codeCompile *lua.FunctionProto) (*lua.LState, error) {
	compL := lua.NewState()
	compL.PreloadModule("Common", LuaModuleLoader)
	if err := ins.doCompileFunc(compL, codeCompile); err != nil {
		return nil, fmt.Errorf("doCompileFunc err:%w", err)
	}
	return compL, nil
}
