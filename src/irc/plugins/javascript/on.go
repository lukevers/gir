/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package javascript

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"github.com/thoj/go-ircevent"
)

// The on function in Javascript should be called as follows:
//
// on("PRIVMSG", function(event) {
//   // stuff
// });
func (p *JavascriptPlugin) on(call otto.FunctionCall) otto.Value {
	event, _ := call.Argument(0).ToString()
	cback, _ := call.Argument(1).ToString()

	p.cb.event = event
	p.cb.id = p.c.Conn.AddCallback(event, func(event *irc.Event) {
		// Only proceed if the channel is correct
		if p.c.Name == event.Arguments[0] {
			e, _ := json.Marshal(struct {
				Message string `json:"message"`
				Channel string `json:"channel"`
				Nick    string `json:"nick"`
				Host    string `json:"host"`
				Source  string `json:"source"`
				User    string `json:"user"`
				Raw     string `json:"raw"`
			}{
				Message: event.Message(),
				Channel: event.Arguments[0],
				Nick:    event.Nick,
				Host:    event.Host,
				Source:  event.Source,
				User:    event.User,
				Raw:     event.Raw,
			})

			p.o.Run(fmt.Sprintf("(%s)(%s);", cback, e))
		}
	})

	return otto.Value{}
}
