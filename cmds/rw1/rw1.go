package main

import (
	// Standard Packages
	"fmt"

	// 3rd Party Packages
	"github.com/robertkrimen/otto"

	// Caltech Packages
	"github.com/caltechlibrary/ostdlib"
)

func main() {
	fmt.Println("Welcome to rw1, the simplest repl built with ostdlib and otto")
	fmt.Println("use .exit to quit the repl, .help to list the dot commands")
	vm := otto.New()
	js := ostdlib.New(vm)
	js.Repl()
}
