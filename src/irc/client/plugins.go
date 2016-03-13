/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package client

import (
	"irc/plugins"
	"irc/plugins/javascript"
	"irc/plugins/lua"
	"path/filepath"
)

func (c *Client) pluginAutoloader() error {
	// Lua
	if files, err := filepath.Glob(lua.GLOB); err != nil {
		return err
	} else {
		for _, channel := range c.channels {
			if channel.Plugins == nil {
				channel.Plugins = make(map[string]map[string]plugins.Plugin)
			} else {
				lua.Unregister(channel)
			}

			channel.Plugins[lua.NAMESPACE] = lua.Register(files, channel)
		}
	}

	// Javascript
	if files, err := filepath.Glob(javascript.GLOB); err != nil {
		return err
	} else {
		for _, channel := range c.channels {
			channel.Plugins[javascript.NAMESPACE] = javascript.Register(files, channel)
		}
	}

	return nil
}
