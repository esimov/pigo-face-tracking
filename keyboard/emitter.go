package keyboard

import (
	"github.com/micmonay/keybd_event"
)

// KeyBonding
type KeyBonding struct {
	keybd_event.KeyBonding
}

// Init initialize a new keybonding event.
func Init() *KeyBonding {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	return &KeyBonding{kb}
}

// TriggerKeypress triggers the key down event for a specific key.
func (kb *KeyBonding) TriggerKeypress(key int) error {
	// Add key to be pressed
	kb.AddKey(key)

	// Press the selected keys
	err := kb.Launching()
	if err != nil {
		return err
	}

	kb.Press()
	kb.Release()

	return nil
}

// Release releases all the registered key events and clears the current instance.
func (kb *KeyBonding) Release() {
	kb.Clear()
}
