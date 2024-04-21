package module

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gookit/color"
)

func FullWidthMessage(message string, color color.Color) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		color.Println("Error getting terminal size:", err)
		return
	}

	var rows, cols int
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)

	messageLen := len(message)

	// Calculate the number of "-" characters needed on each side of the message
	dashLen := (cols - messageLen) / 2

	// Print "-" characters on the left side
	for i := 0; i < dashLen; i++ {
		color.Print("-")
	}

	// Print the message
	color.Print(message)

	// Print "-" characters on the right side
	for i := 0; i < dashLen; i++ {
		color.Print("-")
	}

	// If the message length is odd, print an extra "-" on the right side
	if (cols-messageLen)%2 != 0 {
		color.Print("-")
	}

	color.Println()
}
