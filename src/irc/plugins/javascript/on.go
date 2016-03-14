/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package javascript

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"github.com/thoj/go-ircevent"
	"strings"
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
		// Check to see if the event is occuring in a channel or not
		ch := strings.HasPrefix(event.Arguments[0], "#")
		if !ch {
			// If the event is not in a channel, we'll want to set the "channel"
			// to the buffer that sent the message.
			event.Arguments[0] = event.Nick
		}

		if p.c.Name == event.Arguments[0] || !ch {
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
