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
local function stuff(event)
    -- do stuff here
end

on("PRIVMSG", stuff)

-- or pass an anonymous function
on("PRIVMSG", function(event)
    -- do stuff here
end)
```

```javascript
// javascript
on("PRIVMSG", stuff);
function stuff(event) {
    // do stuff here
};

// or pass an anonymous function
on("PRIVMSG", function(event) {
    // do stuff here
});
```

As you notice, the callback for the `on` function has one parameter which contains an event object. In Lua, `event` is a table of strings, and in JavaScript, `event` is an object. The event that is passed contains information about the event. Here's a JSON example of a `PRIVMSG` sent to the channel `#cs`:

```json
{
    "message": "hello",
    "channel": "#cs",
    "nick":    "lukevers",
    "host":    "localhost",
    "source":  "lukevers!~lukevers@localhost",
    "user":    "~lukevers",
    "raw":     ":lukevers!~lukevers@localhost PRIVMSG #cs :hello"
}
```

#### say

Another important global function is the `say` function. Here's an example using the event passed by the `on` function:

```lua
-- lua
say(event["channel"], "message")
```

```javascript
// javascript
say(event.channel, "message");
```

### Global Objects

Not only are there global functions that exist, but there are also global objects.

#### storage

The storage object allows plugins to interact with your [boltdb](https://github.com/boltdb/bolt) database.

##### use

If the function `use` is never called, the default bucket `"default"` will be used. You can change buckets at any time. To go back to the default bucket, just use the bucket `"default"`. If there was a problem, the function will return `false`, otherwise it will return `true`.

```lua
-- lua
storage:use("bucket")
```

```javascript
// javascript
storage.use("bucket");
```

##### put

If the key already exists, the current value will be overwritten with the new value. If there was a problem, the function will return `false`, otherwise it will return `true`.

```lua
-- lua
storage:put("key", "value")
```

```javascript
// javascript
storage.put("key", "value");
```

##### get

If the key does not exist, the return value will be `nil` for Lua and `null` for JavaScript.

```lua
-- lua
storage::get("key")
```

```javascript
// javascript
storage.get("key");
```

##### remove

If the key does not exist, the return value will be `false`, otherwise it will be `true`.

```lua
-- lua
storage::remove("key")
```

```javascript
// javascript
storage.remove("key");
```
