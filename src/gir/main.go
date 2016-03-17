/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"irc/client"
	"irc/server"
	"log"
	"storage"
	"time"
)

func main() {
	// Setup storage
	store, _ := storage.New("gir.db")

	// Setup bot
	bot := client.New(client.Client{
		Nick:        "gir",
		User:        "gir",
		QuitMessage: "gir",
		Storage:     store,
		Server: &server.Server{
			Host: "localhost",
			Port: 6667,
			Ssl:  false,
		},
	})

	bot.Connect()
	err := bot.Join("#cs", "")
	if err != nil {
		log.Println(err)
	}

	// Enable plugins
	bot.EnablePlugins(30 * time.Second)

	for {
	}
}
