# gir

Another scriptable IRC bot written in Go.

## Plugins / Scripts

Plugins can be written in both JavaScript and Lua. See the [plugins](/plugins) directory for examples.

### Global Functions

There are key global functions that exist.

#### on

The most important global function is the `on` function. Each plugin can call the `on` function as many times as it wishes. The first parameter is the type of event (`PRIVMSG`, `001`, `PING`, or others as defined in [RFC-1459](http://tools.ietf.org/html/rfc1459)), and the second is the function to be called upon that event occuring.

```lua
-- lua
on("PRIVMSG", function(event)
    -- do stuff here
end)
```

```javascript
// javascript
on("PRIVMSG", function(event) {
    // do stuff here
})
```

As you notice, the callback for the `on` function has one parameter which contains an event object. In Lua, `event` is a table of strings, and in JavaScript, `event` is an object. The event that is passed contains information about the event. Here's a JSON example of a `PRIVMSG` sent to the channel `#cs`:

```json
{
    'message': 'hello',
    'channel': '#cs',
    'nick':    'lukevers',
    'host':    'localhost',
    'source':  'lukevers!~lukevers@localhost',
    'user':    '~lukevers',
    'raw':     ':lukevers!~lukevers@localhost PRIVMSG #cs :hello',
}
```

#### say

Another important global function is the `say` function.
