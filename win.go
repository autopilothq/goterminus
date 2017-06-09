// +build windows

package term

import (
	"os"
	"os/exec"
)

// Clear clears all contents from the terminal window
func Clear() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// CursorUp moves the cursor up the desired amount, relative to the current
// position
func CursorUp(amount int) {
	panic("The Banks CLI does not support Windows yet :-(")
}

// CursorDown moves the cursor down the desired amount, relative to the current
// position
func CursorDown(amount int) {
	panic("The Banks CLI does not support Windows yet :-(")
}

// CursorForward moves the cursor forward the desired amount, relative to
// the current position
func CursorForward(amount int) {
	panic("The Banks CLI does not support Windows yet :-(")
}

// CursorBackward moves the cursor backward the desired amount, relative
// to the current position
func CursorBackward(amount int) {
	panic("The Banks CLI does not support Windows yet :-(")
}

// EraseInLine ...
func EraseInLine() {
	panic("The Banks CLI does not support Windows yet :-(")
}
