# Proof of concept for options in Golang

This only works with booleans :D

import this in your project:

```
import "github.com/mimoo/Bool"
```

then you can use the `Bool.Option` type. It contains either a `bool` or an `error`:

```
func thing(arg bool) Bool.Option {
	if arg {
		return Bool.Ok(true)
	}
	return Bool.Err(fmt.Errorf("nope"))
}
```

and you are forced to handle it before you can use the value:

```
switch val := Bool.Match(thing(true)).(type) {
case bool:
    fmt.Println("a bool!", val)
case error:
    fmt.Println("an err!", val)
}
```