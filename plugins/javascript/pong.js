on("PRIVMSG", function(event) {
    console.log(JSON.stringify(event));
    if (event.message.substring(0, 5) === "@pong") {
        say(event.channel, "ping");
    }
});

on("PING", function(event) {
    console.log(1);
    console.log(JSON.stringify(event));
});

on("001", function(event) {
    console.log(3);
    console.log(JSON.stringify(event));
});
