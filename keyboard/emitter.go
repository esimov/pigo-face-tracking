package keyboard

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

func EmitKeyboardPress(key int) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// Select keys to be pressed
	kb.SetKeys(key)

	// Set shift to be pressed
	kb.HasSHIFT(false)

	// Press the selected keys
	err = kb.Launching()
	if err != nil {
		panic(err)
	}

	kb.Press()
	kb.Release()
}
