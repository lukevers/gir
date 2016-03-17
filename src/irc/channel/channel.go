/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package channel

import (
	"github.com/thoj/go-ircevent"
	"irc/plugins"
	"storage"
)

type Channel struct {
	Conn     *irc.Connection
	Name     string
	Password string
	Plugins  map[string]map[string]plugins.Plugin
	Storage  *storage.Storage
}
