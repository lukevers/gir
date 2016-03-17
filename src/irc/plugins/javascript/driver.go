/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package javascript

import (
	"github.com/robertkrimen/otto"
	"irc/channel"
	"irc/plugins"
	"log"
)

const (
	GLOB      = "plugins/javascript/*.js"
	NAMESPACE = "JAVASCRIPT"
)

type JavascriptPlugin struct {
	plugins.Plugin

	c *channel.Channel
	o *otto.Otto

	cb struct {
		event string
		id    int
	}
}

func Register(files []string, c *channel.Channel) map[string]plugins.Plugin {
	plugins := make(map[string]plugins.Plugin)
	for _, file := range files {
		p := &JavascriptPlugin{
			c: c,
			o: otto.New(),
		}

		p.setupAPI()

		script, err := p.o.Compile(file, nil)
		if err != nil {
			log.Println(err)
		}

		p.o.Run(script)
	}

	return plugins
}

func Unregister(c *channel.Channel) {
	for _, plugin := range c.Plugins[NAMESPACE] {
		c.Conn.RemoveCallback(plugin.(*JavascriptPlugin).cb.event, plugin.(*JavascriptPlugin).cb.id)
	}
}

func (p *JavascriptPlugin) setupAPI() {
	p.o.Set("on", p.on)
	p.o.Set("say", p.say)

	storage, _ := p.o.Object(`storage = {}`)
	storage.Set("use", p.storageUse)
	storage.Set("put", p.storagePut)
	storage.Set("get", p.storageGet)
	storage.Set("remove", p.storageRemove)
}
