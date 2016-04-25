package main

import (
	// Standard Packages

	"encoding/json"
	"fmt"
	"log"

	// 3rd Party Packages
	"github.com/robertkrimen/otto"
	"github.com/rsdoiel/replweather"

	// Caltech Packages
	"github.com/caltechlibrary/ostdlib"
)

func main() {
	fmt.Println("Welcome to rw3 adds an NWS object and other basic objects plus help and autocomplete")
	fmt.Println("use .exit to quit the repl, .help to list the dot commands")
	vm := otto.New()
	js := ostdlib.New(vm)
	js.AddExtensions()
	js.AddHelp()

	// Now integrate our replweather parts

	// Define a general purpose error object (This is a convienence, not required)
	errorObject := func(obj *otto.Object, msg string) otto.Value {
		if obj == nil {
			obj, _ = vm.Object(`({})`)
		}
		log.Println(msg)
		obj.Set("status", "error")
		obj.Set("error", msg)
		return obj.Value()
	}

	// Create a NWS object in the JS VM
	nwsObj, _ := js.VM.Object(`NWS = {}`)
	// Add methods to the object
	nwsObj.Set("getRSS", func(call otto.FunctionCall) otto.Value {
		nws, err := replweather.GetNWSRSS()
		if err != nil {
			return errorObject(nil, fmt.Sprintf("NWS.getRSS() failed %s, %s", call.CallerLocation(), err))
		}
		src, err := json.Marshal(nws)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("NWS.getRSS() failed %s, %s", call.CallerLocation(), err))
		}
		// I am creating an object to return from the JSON version of the forcasts
		obj, _ := js.VM.Object(fmt.Sprintf(`result = (%s)`, src))
		// Use js.VM.ToValue() to create the returned value.
		result, err := js.VM.ToValue(obj)
		if err != nil {
			return errorObject(nil, fmt.Sprintf("NWS.getRSS() failed %s, %s", call.CallerLocation(), err))
		}
		return result
	})
	// Add help describing the object which also populates autocomplete, data structure
	js.SetHelp("NWS", "getRSS", []string{}, "Get the RSS feed from weather.gov and return an array of forcasts")

	// Now that all the help has been defined add autocomplete
	js.AddAutoComplete()
	// Run the repl
	js.Repl()
}
