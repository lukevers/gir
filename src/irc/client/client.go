/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package client

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"irc/channel"
	"irc/server"
	"storage"
	"log"
	"time"
)

type Client struct {
	Nick        string
	User        string
	QuitMessage string
	Server      *server.Server
	Storage     *storage.Storage

	channels map[string]*channel.Channel
	conn     *irc.Connection
}

// Create a new irc client
func New(c Client) *Client {
	c.channels = make(map[string]*channel.Channel)
	c.conn = irc.IRC(c.Nick, c.User)
	c.conn.PingFreq = 1 * time.Minute
	c.conn.QuitMessage = c.QuitMessage

	return &c
}

// Connection the client to the IRC server
func (c *Client) Connect() {
	c.conn.Connect(fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port))
}

// Disconnect the client from the IRC server
func (c *Client) Disconnect() {
	c.conn.Quit()
}

// Join a channel
func (c *Client) Join(ch, pass string) error {
	// Check if the client is in the channel
	if _, exists := c.channels[ch]; exists {
		return fmt.Errorf("The client is already in the channel %s", ch)
	}

	// Join channel
	c.conn.Join(fmt.Sprintf("%s %s", ch, pass))

	// Add to channel list
	c.channels[ch] = &channel.Channel{
		Name:     ch,
		Password: pass,
		Conn:     c.conn,
		Storage:  c.Storage,
	}

	return nil
}

// Part a channel
func (c *Client) Part(ch string) error {
	// Check if client is in the channel
	if _, exists := c.channels[ch]; !exists {
		return fmt.Errorf("The client is not currently in the channel %s", ch)
	}

	// Remove from channel list
	delete(c.channels, ch)

	// Leave channel
	c.conn.Part(ch)

	return nil
}

func (c *Client) EnablePlugins(refresh time.Duration) {
	go (func(c *Client) {
		for {
			err := c.pluginAutoloader()
			if err != nil {
				log.Println("[ERROR] %s", err)
			}

			time.Sleep(refresh)
		}
	})(c)
}
