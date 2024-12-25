package main

import (
	"modol-app/helpers"
)

func main() {
	helpers.ConfirmationScreen("Welcome to the Application")

	helpers.ClearScreen()

	helpers.ConfirmationScreen("This is the second screen", "It has multiple messages")
}