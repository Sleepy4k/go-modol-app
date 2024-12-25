//=================================================================================
// Modular Online Learning (MODOL)
//--------------------------------------------------------------------------------
// MODOL is an e-learning application designed to provide a modular
// learning experience, where each module can stand alone or be integrated with
// other modules. The application is built using the Go programming language for
// efficiency and high performance.
//--------------------------------------------------------------------------------
// Version: 1.0.0
// Date: 2024-12-25
// License: Unlicensed
// Environment: Console
// OS: Windows
// Language: Go
//=================================================================================

package main

import (
	"fmt"
	"modol-app/helpers"
)

func main() {
  var choice int

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    helpers.DisplayMenu()

    for choice < 1 || choice > 2 {
      fmt.Print("Enter your choice: ")
      fmt.Scanln(&choice)
    }

    switch choice {
    case 1:
      fmt.Println("Starting learning...")
      break
    case 2:
      isRunning = true
      helpers.ClearScreen()
      fmt.Println("Thank you for using MODOL. Goodbye!")  
      break
    default:
      fmt.Println("Invalid choice. Please try again.")
      helpers.ConfirmationScreen()
      break
    }

    choice = 0
  } 
}