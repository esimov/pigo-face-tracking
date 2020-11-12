package conn

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/esimov/pigo-face-tracking/keyboard"
	"github.com/gorilla/websocket"
	"github.com/micmonay/keybd_event"
)

// HttpParams http/websocket connection parameters
type HttpParams struct {
	Address string
	Prefix  string
	Root    string
}

// wsclient is a middleman between the websocket connection and the hub.
type wsclient struct {
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// socketChan is a channel used for sending the detection results through
var socketChan = make(chan string)

var (
	kb keyboard.KeyBonding
	ws *wsclient
)

// A server application calls the Upgrade method from an HTTP request handler to initiate a connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Init initializes the webserver and websocket connection
func Init(p *HttpParams) {
	var err error
	p.Root, err = filepath.Abs(p.Root)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("serving %s as %s on %s", p.Root, p.Prefix, p.Address)
	http.Handle(p.Prefix, http.StripPrefix(p.Prefix, http.FileServer(http.Dir(p.Root))))
	http.HandleFunc("/ws", wsHandler)

	go ws.run()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr + " " + r.Method + " " + r.URL.String())
		http.DefaultServeMux.ServeHTTP(w, r)
	})
	httpServer := http.Server{
		Addr:    p.Address,
		Handler: handler,
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

	kb = *keyboard.Init()
}

// wsHandler is the websocket connection handler
func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Upgrade the http connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	ws = &wsclient{conn: conn}
	go ws.readSocket()
}

// readSocket listen for new messages being sent to the websocket
func (ws *wsclient) readSocket() {
	defer func() {
		ws.conn.Close()
	}()

	for {
		messageType, msg, err := ws.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return
		}
		socketChan <- string(msg)

		if err := ws.conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

// run listens on the opened websocket connection and recieve the detection results concurrently.
func (ws *wsclient) run() {
	var cmd int
	defer func() {
		close(socketChan)
	}()

	for {
		select {
		case key, ok := <-socketChan:
			if ok {
				// Check if the connection is open and the channel is not closed.
				switch key {
				case "down":
					cmd = keybd_event.VK_DOWN
					break
				case "up":
					cmd = keybd_event.VK_UP
					break
				case "right":
					cmd = keybd_event.VK_RIGHT
					break
				case "left":
					cmd = keybd_event.VK_LEFT
					break
				}
				kb.TriggerKeypress(cmd)
			} else {
				// Release the key bindings.
				kb.Release()
				// The hub closed the channel.
				ws.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
		}
	}
}
