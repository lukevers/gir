/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package lua

import (
	Lua "github.com/yuin/gopher-lua"
)

func (p *LuaPlugin) storageUse(state *Lua.LState) int {
	table := state.ToString(1)

	err := p.c.Storage.Use(table)
	if err != nil {
		p.l.Push(Lua.LFalse)
		return 1
	}

	p.l.Push(Lua.LTrue)
	return 1
}

func (p *LuaPlugin) storagePut(state *Lua.LState) int {
	key := state.ToString(1)
	val := state.ToString(2)

	err := p.c.Storage.Put(key, val)
	if err != nil {
		p.l.Push(Lua.LFalse)
		return 1
	}

	p.l.Push(Lua.LTrue)
	return 1
}

func (p *LuaPlugin) storageGet(state *Lua.LState) int {
	key := state.ToString(1)

	val, err := p.c.Storage.Get(key)
	if err != nil {
		p.l.Push(Lua.LNil)
		return 1
	}

	p.l.Push(Lua.LString(val))
	return 1
}

func (p *LuaPlugin) storageRemove(state *Lua.LState) int {
	key := state.ToString(1)

	err := p.c.Storage.Remove(key)
	if err != nil {
		p.l.Push(Lua.LFalse)
		return 1
	}

	p.l.Push(Lua.LTrue)
	return 1
}
