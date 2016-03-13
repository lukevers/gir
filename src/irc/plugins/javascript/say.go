/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package javascript

import (
	"github.com/robertkrimen/otto"
)

// The say function in Javascript should be called as follows:
//
// say("#channel", "message")
//
// or with the global event variable:
//
// say(event.channel, "message")
func (p *JavascriptPlugin) say(call otto.FunctionCall) otto.Value {
	channel, _ := call.Argument(0).ToString()
	message, _ := call.Argument(1).ToString()

	p.c.Conn.Privmsg(channel, message)
	return otto.Value{}
}
