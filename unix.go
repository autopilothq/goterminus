// +build !windows

package term

import (
	"fmt"
	"os"
	"os/exec"
)

// Clear clears all contents from the terminal window
func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// CursorUp moves the cursor up the desired amount, relative to the current
// position
func CursorUp(amount int) {
	fmt.Printf("\x1b[%dA", amount)
}

// CursorDown moves the cursor down the desired amount, relative to the current
// position
func CursorDown(amount int) {
	fmt.Printf("\x1b[%dB", amount)
}

// CursorForward moves the cursor forward the desired amount, relative to
// the current position
func CursorForward(amount int) {
	fmt.Printf("\033[%dC", amount)
}

// CursorBackward moves the cursor backward the desired amount, relative
// to the current position
func CursorBackward(amount int) {
	fmt.Printf("\033[%dD", amount)
}

// EraseInLine ...
func EraseInLine() {
	fmt.Print("\x1b[0K")
}
