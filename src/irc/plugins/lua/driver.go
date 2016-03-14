/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package lua

import (
	Lua "github.com/yuin/gopher-lua"
	"irc/channel"
	"irc/plugins"
	"log"
	"strings"
)

const (
	GLOB      = "plugins/lua/*.lua"
	NAMESPACE = "LUA"
)

type LuaPlugin struct {
	plugins.Plugin

	c *channel.Channel
	l *Lua.LState

	cb struct {
		event string
		id    int
	}
}

func Register(files []string, c *channel.Channel) map[string]plugins.Plugin {
	plugins := make(map[string]plugins.Plugin)
	for _, file := range files {
		p := &LuaPlugin{
			l: Lua.NewState(Lua.Options{}),
			c: c,
		}

		p.setupAPI()

		if err := p.l.DoFile(file); err != nil {
			log.Println("Erorr running plugin from file: ", err)
		} else {
			plugins[file] = p
		}
	}

	return plugins
}

func Unregister(c *channel.Channel) {
	for _, plugin := range c.Plugins[NAMESPACE] {
		c.Conn.RemoveCallback(plugin.(*LuaPlugin).cb.event, plugin.(*LuaPlugin).cb.id)
	}
}

func (p *LuaPlugin) setupAPI() {
	p.l.SetGlobal("on", p.l.NewFunction(p.on))
	p.l.SetGlobal("say", p.l.NewFunction(p.say))
}
