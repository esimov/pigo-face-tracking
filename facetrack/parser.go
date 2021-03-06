package facetrack

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"syscall/js"
	"time"
)

// Detector struct holds the main components of the fetching operation.
type Detector struct {
	window js.Value
}

// NewDetector initializes a new constructor function.
func NewDetector() *Detector {
	var d Detector
	d.window = js.Global()

	return &d
}

// ParseCascade loads and parse the cascade file through the
// Javascript `location.href` method, using the `js/syscall` package.
// It will return the cascade file encoded into a byte array.
func (d *Detector) ParseCascade(path string) ([]byte, error) {
	href := js.Global().Get("location").Get("href")
	u, err := url.Parse(href.String())
	if err != nil {
		return nil, err
	}
	u.Path = path
	u.RawQuery = fmt.Sprint(time.Now().UnixNano())

	resp, err := http.Get(u.String())
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%v cascade file is missing", u.String()))
	}
	defer resp.Body.Close()

	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	uint8Array := js.Global().Get("Uint8Array").New(len(buffer))
	js.CopyBytesToJS(uint8Array, buffer)

	jsbuf := make([]byte, uint8Array.Get("length").Int())
	js.CopyBytesToGo(jsbuf, uint8Array)

	return jsbuf, nil
}

// Log calls the `console.log` Javascript function
func (d *Detector) Log(args ...interface{}) {
	d.window.Get("console").Call("log", args...)
}
