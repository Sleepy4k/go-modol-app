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
  "modol-app/handlers/auth"
  ctrls "modol-app/controllers"
)

func main() {
  var choice int

  for isLogged := false; !isLogged; {
    helpers.ClearScreen()
    helpers.DisplayAuthMenu()

    for choice < 1 || choice > 3 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    switch choice {
    case 1:
      isLogged = ctrls.AuthController{}.Login()
    case 2:
      ctrls.AuthController{}.Register()
    case 3:
      isLogged = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
      helpers.ConfirmationScreen()
    }

    if choice != 3 {
      choice = 0
    }
  }

  if choice == 3 {
    helpers.ClearScreen()
    fmt.Println("Terima kasih telah menggunakan MODOL. Selamat tinggal!")
    return
  }

  if !auth.IsAuthenticated(auth.CurrentUser) {
    fmt.Println("Anda harus masuk terlebih dahulu.")
    helpers.ConfirmationScreen()
    return
  }

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    helpers.DisplayMainMenu()

    for choice < 1 || choice > 5 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    switch choice {
    case 1:
      fmt.Println("Starting learning...")
    case 2:
      fmt.Println("Starting classes...")
    case 3:
      fmt.Println("Starting grades...")
    case 4:
      ctrls.SettingController{}.Index()
    case 5:
      isRunning = true
      helpers.ClearScreen()
      fmt.Println("Terima kasih telah menggunakan MODOL. Selamat tinggal!")
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
      helpers.ConfirmationScreen()
    }

    if choice != 5 {
      choice = 0
    }
  } 
}