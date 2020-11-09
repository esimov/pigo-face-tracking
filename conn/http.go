package conn

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/esimov/pigo-face-tracking/keyboard"
	"github.com/gorilla/websocket"
)

// HttpParams http/websocket connection parameters
type HttpParams struct {
	Address string
	Prefix  string
	Root    string
}

// socketChan is a channel used for sending the detection results through
var socketChan = make(chan string)

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
	go run()

	log.Printf("serving %s as %s on %s", p.Root, p.Prefix, p.Address)
	http.Handle(p.Prefix, http.StripPrefix(p.Prefix, http.FileServer(http.Dir(p.Root))))
	http.HandleFunc("/ws", wsHandler)

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
}

// readSocket listen for new messages being sent to the websocket
func readSocket(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return
		}
		socketChan <- string(msg)

		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return
		}
	}
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
	go readSocket(conn)
}

// Listen on the opened websocket connection and recieve the detection results concurrently.
func run() {
	defer func() {
		close(socketChan)
	}()

	for {
		select {
		case det, ok := <-socketChan:
			if ok {
				detections := strings.Split(det, ",")
				keyboard.EmitKeyboardPress()
				fmt.Println(detections)
			}
		}
	}
}
