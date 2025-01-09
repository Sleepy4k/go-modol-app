package controllers

import (
	"fmt"
	"modol-app/enums"
	"modol-app/handlers/auth"
	"modol-app/helpers"
	"modol-app/models/class"
  "modol-app/models/course_has_class"
  "modol-app/models/quiz"
  "modol-app/models/assignment"
  "modol-app/models/discuss"
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

    for choice < 1 || choice > 6 {
      helpers.ClearScreen()
      helpers.DisplaySettingMenu()

      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    if auth.IsSiswa(auth.CurrentUser) && (choice > 1 && choice < 6) {
      fmt.Println("Anda tidak memiliki izin untuk mengakses menu ini.")
      choice = 0
      helpers.ConfirmationScreen()
      continue
    }

    switch choice {
    case 1:
      sc.Profile()
    case 2:
      sc.ManageClass()
    case 3:
      sc.ManageAssignment()
    case 4:
      sc.ManageQuiz()
    case 5:
      sc.ManageDiscussion()
    case 6:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 6 {
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

    for choice < 1 || choice > 5 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    switch choice {
    case 1:
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

      clasess := class.ListClasses(sortType)

      helpers.ClearScreen()

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
      var data class.Class

      fmt.Print("ID Kelas: ")
      fmt.Scanln(&ID)

      currClass := class.FindClassByID(ID)

      if currClass.ID == 0 {
        fmt.Println("Kelas tidak ditemukan.")
      } else {
        fmt.Print("Nama Kelas: ")
        data.Nama = helpers.GetInlineInput()

        class.UpdateClass(ID, data)
      }
    case 4:
      var ID int

      fmt.Print("ID Kelas: ")
      fmt.Scanln(&ID)

      class.DeleteClass(ID)
    case 5:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 5 {
      choice = 0
      helpers.ConfirmationScreen()
    }
  }
}

/**
  * Displays the assignment management settings
  *
  * @return void
  */
func (sc SettingController) ManageAssignment() {
  var choice int

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    helpers.DisplayManageAssignmentMenu()

    for choice < 1 || choice > 5 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    switch choice {
    case 1:
      assignments := assignment.GetAll()

      helpers.ClearScreen()

      fmt.Println("=== Daftar Tugas ===")
      for _, assignment := range assignments {
        fmt.Printf("%d. %s (%s)\n", assignment.ID, assignment.Title, assignment.Description)
      }
    case 2:
      var data assignment.Assignment

      fmt.Print("Judul Tugas: ")
      data.Title = helpers.GetInlineInput()

      fmt.Print("Deskripsi Tugas: ")
      data.Description = helpers.GetInlineInput()

      fmt.Print("ID Kursus: ")
      fmt.Scanln(&data.CourseID)

      course := course_has_class.ListClassesAssignedToCourse(data.CourseID)

      if len(course) == 0 {
        fmt.Println("Kursus tidak ditemukan.")
      } else {
        if course[0].ID == 0 {
          fmt.Println("Kursus tidak ditemukan.")
        } else {
          data.ClassID = course[0].ID

          assignment.InsertAssignment(data)

          fmt.Println("Tugas berhasil ditambahkan.")
        }
      }
    case 3:
      var ID int

      fmt.Print("ID Tugas: ")
      fmt.Scanln(&ID)

      currQuiz := quiz.FindQuizByID(ID)

      if currQuiz.ID == 0 {
        fmt.Println("Tugas tidak ditemukan.")
      } else {
        fmt.Print("Judul Tugas: ")
        currQuiz.Title = helpers.GetInlineInput()

        fmt.Print("Deskripsi Tugas: ")
        currQuiz.Description = helpers.GetInlineInput()

        quiz.UpdateQuizByID(ID, currQuiz)

        fmt.Println("Tugas berhasil diperbarui.")
      }
    case 4:
      var ID int

      fmt.Print("ID Tugas: ")
      fmt.Scanln(&ID)

      quiz.DeleteQuizByID(ID)

      fmt.Println("Tugas berhasil dihapus.")
    case 5:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 5 {
      choice = 0
      helpers.ConfirmationScreen()
    }
  }
}

/**
  * Displays the quiz management settings
  *
  * @return void
  */
func (sc SettingController) ManageQuiz() {
  var choice int

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    helpers.DisplayManageQuizMenu()

    for choice < 1 || choice > 5 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    switch choice {
    case 1:
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

      quizzes := quiz.ListQuizzes(sortType)

      helpers.ClearScreen()

      fmt.Println("=== Daftar Kuis ===")
      for _, quiz := range quizzes {
        fmt.Printf("%d. %s\n", quiz.ID, quiz.Title)
      }
    case 2:
      var data quiz.Quiz

      fmt.Print("Judul Kuis: ")
      data.Title = helpers.GetInlineInput()

      fmt.Print("Deskripsi Kuis: ")
      data.Description = helpers.GetInlineInput()

      quiz.InsertQuiz(data)
    case 3:
      var ID int
      var data quiz.Quiz

      fmt.Print("ID Kuis: ")
      fmt.Scanln(&ID)

      currQuiz := quiz.FindQuizByID(ID)

      if currQuiz.ID == 0 {
        fmt.Println("Kuis tidak ditemukan.")
      } else {
        fmt.Print("Judul Kuis: ")
        data.Title = helpers.GetInlineInput()

        fmt.Print("Deskripsi Kuis: ")
        data.Description = helpers.GetInlineInput()

        quiz.UpdateQuizByID(ID, data)
      }
    case 4:
      var ID int

      fmt.Print("ID Kuis: ")
      fmt.Scanln(&ID)

      quiz.DeleteQuizByID(ID)
    case 5:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 5 {
      choice = 0
      helpers.ConfirmationScreen()
    }
  }
}

/**
  * Displays the discussion management settings
  *
  * @return void
  */
func (sc SettingController) ManageDiscussion() {
  var choice int

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    helpers.DisplayManageDiscussionMenu()

    for choice < 1 || choice > 5 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    switch choice {
    case 1:
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

      discussions := discuss.ListDiscusses(sortType)

      helpers.ClearScreen()

      fmt.Println("=== Daftar Diskusi ===")
      for _, discussion := range discussions {
        fmt.Printf("%d. %s\n", discussion.ID, discussion.Content)
      }
    case 2:
      var data discuss.Discuss

      fmt.Print("Judul Diskusi: ")
      data.Content = helpers.GetInlineInput()

      discuss.InsertDiscuss(data)
    case 3:
      var ID int
      var data discuss.Discuss

      fmt.Print("ID Diskusi: ")
      fmt.Scanln(&ID)

      currDiscussion := discuss.FindDiscussByID(ID)

      if currDiscussion.ID == 0 {
        fmt.Println("Diskusi tidak ditemukan.")
      } else {
        fmt.Print("Judul Diskusi: ")
        data.Content = helpers.GetInlineInput()

        discuss.UpdateDiscussionByID(ID, data)
      }
    case 4:
      var ID int

      fmt.Print("ID Diskusi: ")
      fmt.Scanln(&ID)

      discuss.DeleteDiscussById(ID)
    case 5:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 5 {
      choice = 0
      helpers.ConfirmationScreen()
    }
  }
}