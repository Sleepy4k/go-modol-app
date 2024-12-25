package helpers

import "fmt"

/**
 * Prints the header for the application
 *
 * @return void
 */
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

/**
 * Prints the confirmation screen with the given messages
 *
 * @return void
 */
func ConfirmationScreen(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}

	if len(messages) != 0 {
		fmt.Println()
	}

	fmt.Println("Press Enter to Continue")
	fmt.Scanln()
}