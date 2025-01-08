package controllers

import (
	"fmt"
	"modol-app/enums"
	"modol-app/handlers/auth"
	"modol-app/helpers"
	"modol-app/models/class"
)

/**
 * SettingController struct
 */
type SettingController struct {}

/**
  * Displays the setting menu
  *
  * @return void
  */
func (sc SettingController) Index() {
  for isRunning := false; !isRunning; {
    var choice int

    for choice < 1 || choice > 3 {
      helpers.ClearScreen()
      helpers.DisplaySettingMenu()

      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    switch choice {
    case 1:
      sc.Profile()
    case 2:
      if !auth.IsSiswa(auth.CurrentUser) {
        sc.ManageClass()
      } else {
        fmt.Println("Anda tidak memiliki izin untuk mengakses menu ini.")
      }
    case 3:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 3 {
      helpers.ConfirmationScreen()
    }
  }
}

/**
  * Displays the profile settings
  *
  * @return void
  */
func (sc SettingController) Profile() {
  user := auth.CurrentUser

  fmt.Println("=== Profil Akun ===")
  fmt.Printf("Nama: %s\n", user.Nama)
  fmt.Printf("NIM: %s\n", user.NIM)
  fmt.Printf("Role: %s\n", user.Role)
}

/**
  * Displays the class management settings
  *
  * @return void
  */
func (sc SettingController) ManageClass() {
  var choice int

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    helpers.DisplayManageClassMenu()

    for choice < 1 || choice > 4 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    switch choice {
    case 1:
      sortType := enums.ASC

      fmt.Println("1. Urutkan berdasarkan nama (A-Z)")
      fmt.Println("2. Urutkan berdasarkan nama (Z-A)")

      for sortType < 1 || sortType > 2 {
        fmt.Print("Masukan pilihan anda: ")
        fmt.Scanln(&sortType)
      }

      if sortType == 2 {
        sortType = enums.DESC
      }

      clasess := class.ListClasses(sortType)

      fmt.Println("=== Daftar Kelas ===")
      for _, class := range clasess {
        fmt.Printf("%d. %s\n", class.ID, class.Nama)
      }
    case 2:
      var data class.Class

      fmt.Print("Nama Kelas: ")
      data.Nama = helpers.GetInlineInput()

      class.InsertClass(data)
    case 3:
      var ID int

      fmt.Print("ID Kelas: ")
      fmt.Scanln(&ID)

      class.DeleteClass(ID)
    case 4:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 4 {
      choice = 0
      helpers.ConfirmationScreen()
    }
  }
}