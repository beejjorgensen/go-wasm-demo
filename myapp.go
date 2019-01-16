package main

import "syscall/js"

func buttonCallback(args []js.Value) {
	document := js.Global().Get("document")
	box := document.Call("getElementById", "box") // From the DOM
	box.Set("innerHTML", "Button 2 clicked!")
}

func main() {
	println("Hello, world!")

	document := js.Global().Get("document")
	box := document.Call("getElementById", "box") // From the DOM
	box.Set("innerHTML", "Hi there!")

	cb := js.NewCallback(buttonCallback)

	button2 := document.Call("getElementById", "button2")
	button2.Call("addEventListener", "click", cb)

	// Sleep forever
	c := make(chan struct{}, 0)
	<-c
}
