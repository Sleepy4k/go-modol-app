package controllers

import (
  "fmt"
  "modol-app/enums"
  "modol-app/handlers/auth"
  "modol-app/helpers"
  "modol-app/models/course"
  "modol-app/models/class_has_user"
  "modol-app/models/course_has_class"
)

/**
  * HomeController struct
  */
type HomeController struct {}

/**
  * Displays the course menu
  *
  * @return void
  */
func (hc HomeController) Index() {
  for isRunning := false; !isRunning; {
    var choice int

    for choice < 1 || choice > 4 {
      helpers.ClearScreen()
      helpers.DisplayHomeMenu()

      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    switch choice {
    case 1:
      hc.List()
    case 2:
      if auth.IsSiswa(auth.CurrentUser) {
        hc.Show()
      } else {
        fmt.Println("Anda tidak memiliki izin untuk mengakses menu ini.")
      }
    case 3:
      if auth.IsSiswa(auth.CurrentUser) {
        hc.Register()
      } else {
        fmt.Println("Anda tidak memiliki izin untuk mengakses menu ini.")
      }
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

/**
  * Displays the list of courses
  *
  * @return void
  */
func (hc HomeController) List() {
  sortType := enums.ASC
  var choice int

  fmt.Println("1. Urutkan berdasarkan nama (A-Z)")
  fmt.Println("2. Urutkan berdasarkan nama (Z-A)")

  for choice < 1 || choice > 2 {
    fmt.Print("Masukan pilihan anda: ")
    fmt.Scanln(&choice)
  }

  if choice == 2 {
    sortType = enums.DESC
  }

  courses := course.ListCourses(sortType)

  helpers.ClearScreen()

  fmt.Println("=== Daftar Kursus ===")
  for _, course := range courses {
    fmt.Printf("%d. %s\n", course.ID, course.Nama)
  }
}

/**
  * Displays the details of a course
  *
  * @return void
  */
func (hc HomeController) Show() {
  var ID int

  courses := course.ListCourses(enums.ASC)

  fmt.Println("=== Daftar Kursus ===")
  for _, course := range courses {
    fmt.Printf("%d. %s\n", course.ID, course.Nama)
  }

  fmt.Println("===============")
  fmt.Print("ID Kursus: ")
  fmt.Scanln(&ID)

  course := course.FindCourseByID(ID)

  if course == nil || course.ID == 0 {
    fmt.Println("Kelas tidak ditemukan.")
  } else {
    class := course_has_class.ListClassesAssignedToCourse(ID)[0]
    totalUsers := class_has_user.ListUsersAssignedToClass(class.ID)

    fmt.Printf("Nama Kursus: %s\n", course.Nama)
    fmt.Printf("Kelas: %s\n", class.Nama)
    fmt.Printf("Jumlah Siswa: %d\n", len(totalUsers))
  }
}

/**
  * Registers a user to a course
  *
  * @return void
  */
func (hc HomeController) Register() {
  isAlreadyAssigned := class_has_user.ListClassesAssignedToUser(auth.CurrentUser.ID)

  if len(isAlreadyAssigned) > 0 {
    fmt.Println("Anda sudah terdaftar pada kelas.")
    return
  }

  var ID int

  courses := course.ListCourses(enums.ASC)

  fmt.Println("=== Daftar Kursus ===")
  for _, course := range courses {
    fmt.Printf("%d. %s\n", course.ID, course.Nama)
  }

  fmt.Println("===============")
  fmt.Print("ID Kursus: ")
  fmt.Scanln(&ID)

  classID := course_has_class.ListClassesAssignedToCourse(ID)[0].ID

  if classID == 0 {
    fmt.Println("Kelas tidak ditemukan.")
    return
  }

  class_has_user.AssignUserToClass(classID, auth.CurrentUser.ID)
}