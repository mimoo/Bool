package main

import (
	"fmt"

	"github.com/mimoo/Bool"
)

func thing(arg bool) Bool.Option {
	if arg {
		return Bool.Ok(true)
	}
	return Bool.Err(fmt.Errorf("nope"))
}

func main() {
	switch val := Bool.Match(thing(true)).(type) {
	case bool:
		fmt.Println("a bool!", val)
	case error:
		fmt.Println("an err!", val)
	}
}
