package helpers

import (
  "fmt"
  "modol-app/config"
)

/**
  * Prints the header for the application
  *
  * @return void
  */
func ClearScreen() {
  fmt.Print("\033[H\033[2J")
}

/**
  * Prints the authentication menu for the application
  *
  * @return void
  */
func DisplayAuthMenu() {
  fmt.Println("=== Autentikasi ===")
  fmt.Println("1. Masuk")
  fmt.Println("2. Daftar")
  fmt.Println("3. Keluar")
  fmt.Println("=======================================")
}

/**
  * Prints the setting menu for the application
  *
  * @return void
  */
func DisplaySettingMenu() {
  fmt.Println("=== Pengaturan ===")
  fmt.Println("1. Profil")
  fmt.Println("2. Kelola Kelas")
  fmt.Println("3. Keluar")
  fmt.Println("=======================================")
}

/**
  * Prints the manage class menu for the application
  *
  * @return void
  */
func DisplayManageClassMenu() {
  fmt.Println("=== Kelola Kelas ===")
  fmt.Println("1. Lihat Kelas")
  fmt.Println("2. Tambah Kelas")
  fmt.Println("3. Hapus Kelas")
  fmt.Println("4. Kembali")
  fmt.Println("=======================================")
}

/**
  * Prints the main menu for the application
  *
  * @return void
  */
func DisplayMainMenu() {
  fmt.Printf("=== %s ===\n", config.APP_NAME)
  fmt.Println("1. Beranda")
  fmt.Println("2. Kelas")
  fmt.Println("3. Nilai")
  fmt.Println("4. Pengaturan")
  fmt.Println("5. Keluar")
  fmt.Println("=======================================")
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