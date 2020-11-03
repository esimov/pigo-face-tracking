package detector

import (
	"syscall/js"

	"github.com/esimov/pigo-face-tracking/conn"
)

// InitWebSocket initializes the websocket connection.
func (c *Canvas) InitWebSocket() {
	ws := conn.GetParams()
	c.ws = js.Global().Get("WebSocket").New("ws://" + ws.Address + "/ws")
	c.Log("Attempting websocket connection...")

	openCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		c.Log("Websocket connection open!")
		return nil
	})

	closeCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		c.Log("Websocket connection closed: ", event)
		return nil
	})

	errorCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		error := args[0]
		c.Log("Websocket error:", error)
		return nil
	})

	c.ws.Call("addEventListener", "open", openCallback)
	c.ws.Call("addEventListener", "close", closeCallback)
	c.ws.Call("addEventListener", "error", errorCallback)
}

// Send sends the value through the socket.
func (c *Canvas) Send(value string) {
	c.ws.Call("send", value)
}
