on("PRIVMSG", function(event) {
    if (event.message.substring(0, 5) === "@pong") {
        say(event.channel, "ping");
    }
});
