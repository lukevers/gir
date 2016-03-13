/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package lua

import (
	Lua "github.com/yuin/gopher-lua"
)

// The say function in Lua should be called as follows:
//
// say("#channel", "message")
//
// or with the global event variable:
//
// say(event["channel"], "message")
func (p *LuaPlugin) say(state *Lua.LState) int {
	channel := state.ToString(1)
	message := state.ToString(2)
	p.c.Conn.Privmsg(channel, message)
	return 1
}
