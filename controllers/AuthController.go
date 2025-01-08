package controllers

import (
	"fmt"
	"modol-app/enums"
	"modol-app/handlers/auth"
	"modol-app/helpers"
)

/**
 * AuthController struct
 */
type AuthController struct {}

/**
  * Handle the authentication process
  *
  * @return void
  */
func (ac AuthController) Login() bool {
  var NIM, password string
  isLogged := false

  fmt.Println("=== Masuk ===")

  fmt.Print("NIM: ")
  fmt.Scan(&NIM)

  fmt.Print("Password: ")
  fmt.Scan(&password)

  helpers.ClearScreen()
  auth.Authentic(NIM, password)

  if !auth.IsAuthenticated(auth.CurrentUser) {
    fmt.Println("Autentikasi gagal. Silakan coba lagi.")
  } else {
    fmt.Printf("Autentikasi berhasil. Selamat datang, %s!\n", auth.CurrentUser.Nama)
    isLogged = true
  }

  helpers.ConfirmationScreen()

  return isLogged
}

/**
  * Handle the registration process
  *
  * @return void
  */
func (ac AuthController) Register() {
  fmt.Println("=== Daftar ===")

  var NIM, nama, password, confirmPassword string

  fmt.Print("Nama: ")
  nama = helpers.GetInlineInput()

  fmt.Print("NIM: ")
  fmt.Scan(&NIM)

  for isCorrect := false; !isCorrect; {
    fmt.Print("Password: ")
    fmt.Scan(&password)

    fmt.Print("Konfirmasi Password: ")
    fmt.Scan(&confirmPassword)

    if password != confirmPassword {
      fmt.Println("Password tidak sama. Silakan coba lagi.")
    } else {
      isCorrect = true
    }
  }

  helpers.ClearScreen()
  auth.CreateUser(NIM, nama, enums.Siswa, password)
}