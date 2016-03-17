/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package lua

import (
	"github.com/thoj/go-ircevent"
	Lua "github.com/yuin/gopher-lua"
	"log"
	"strings"
)

// The on function in Lua should be called as follows:
//
// on("PRIVMSG", function (event)
//   -- stuff
// end)
func (p *LuaPlugin) on(state *Lua.LState) int {
	event := state.ToString(1)
	cback := state.ToFunction(2)

	p.cb.event = event
	p.cb.id = p.c.Conn.AddCallback(event, func(event *irc.Event) {
		// Check to see if the event is occuring in a channel or not
		ch := strings.HasPrefix(event.Arguments[0], "#")
		if !ch {
			// If the event is not in a channel, we'll want to set the "channel"
			// to the buffer that sent the message.
			event.Arguments[0] = event.Nick
		}

		if p.c.Name == event.Arguments[0] || !ch {
			table := new(Lua.LTable)
			table.RawSetString("message", Lua.LString(event.Message()))
			table.RawSetString("channel", Lua.LString(event.Arguments[0]))
			table.RawSetString("nick", Lua.LString(event.Nick))
			table.RawSetString("host", Lua.LString(event.Host))
			table.RawSetString("source", Lua.LString(event.Source))
			table.RawSetString("user", Lua.LString(event.User))
			table.RawSetString("raw", Lua.LString(event.Raw))

			if err := p.l.CallByParam(Lua.P{
				Fn:      cback,
				NRet:    1,
				Protect: true,
			}, table); err != nil {
				log.Println("Error calling lua function: ", err)
			}
		}
	})

	return 1
}
