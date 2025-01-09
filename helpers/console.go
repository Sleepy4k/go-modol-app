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
  fmt.Println("============= Autentikasi ============")
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
  fmt.Println("============= Pengaturan ============")
  fmt.Println("1. Profil")
  fmt.Println("2. Kelola Kelas")
  fmt.Println("3. Kembali")
  fmt.Println("=======================================")
}

/**
  * Prints the manage class menu for the application
  *
  * @return void
  */
func DisplayManageClassMenu() {
  fmt.Println("============ Kelola Kelas ============")
  fmt.Println("1. Lihat Kelas")
  fmt.Println("2. Tambah Kelas")
  fmt.Println("3. Hapus Kelas")
  fmt.Println("4. Kembali")
  fmt.Println("=======================================")
}

/**
  * Prints the home menu for the application
  *
  * @return void
  */
func DisplayHomeMenu() {
  fmt.Println("=============== Beranda =============")
  fmt.Println("1. List Kursus")
  fmt.Println("2. Detail Kursus")
  fmt.Println("3. Daftar Kursus")
  fmt.Println("4. Kembali")
  fmt.Println("=======================================")
}

/**
  * Prints the class menu for the application
  *
  * @return void
  */
func DisplayClassMenu() {
  fmt.Println("================ Kelas ================")
  fmt.Println("1. Materi")
  fmt.Println("2. Tugas")
  fmt.Println("3. Kuis")
  fmt.Println("4. Forum")
  fmt.Println("5. Kembali")
  fmt.Println("=======================================")
}

/**
  * Prints the forum menu for the application
  *
  * @return void
  */
func DisplayForumMenu() {
  fmt.Println("=============== Forum ================")
  fmt.Println("1. List Diskusi")
  fmt.Println("2. Tambah Diskusi")
  fmt.Println("3. Kembali")
  fmt.Println("=======================================")
}

/**
  * Prints the grade menu for the application
  *
  * @return void
  */
func DisplayGradeMenu() {
  fmt.Println("================ Nilai ================")
  fmt.Println("1. List Nilai")
  fmt.Println("2. Kembali")
  fmt.Println("=======================================")
}

/**
  * Prints the main menu for the application
  *
  * @param string curUsername
  * @return void
  */
func DisplayMainMenu(curUsername string) {
  fmt.Printf("=== %s ===\n", config.APP_NAME)
  fmt.Printf("Hai, %s!\n", curUsername)
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

  if len(messages) > 0 {
    fmt.Println()
  }

  fmt.Println("Tekan tombol enter untuk melanjutkan...")
  fmt.Scanln()
}