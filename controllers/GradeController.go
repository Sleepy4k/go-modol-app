package controllers

import (
  "fmt"
  "modol-app/enums"
  "modol-app/helpers"
  "modol-app/handlers/auth"
  "modol-app/models/quiz"
  "modol-app/models/user_has_quiz"
)

/**
  * GradeController struct
  */
type GradeController struct {}

/**
  * Displays the grade menu
  *
  * @return void
  */
func (gc GradeController) Index() {
  for isRunning := false; !isRunning; {
    var choice int

    for choice < 1 || choice > 2 {
      helpers.ClearScreen()
      helpers.DisplayGradeMenu()

      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    switch choice {
    case 1:
      if auth.IsSiswa(auth.CurrentUser) {
        gc.List()
      } else {
        fmt.Println("Anda tidak memiliki izin untuk mengakses menu ini.")
      }
    case 2:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 2 {
      choice = 0
      helpers.ConfirmationScreen()
    }
  }
}

/**
  * Displays the list of grades
  *
  * @return void
  */
func (gc GradeController) List() {
  fmt.Println("=== Daftar Nilai ===")

  quizes := quiz.ListQuizzes(enums.ASC)
  userQuizes := user_has_quiz.GetAllByUserID(auth.CurrentUser.ID)

  if len(quizes) == 0 {
    fmt.Println("Tidak ada kuis yang tersedia.")
    helpers.ConfirmationScreen()
    return
  }

  for _, quiz := range quizes {
    var isFound bool

    fmt.Printf("Kuis: %s\n", quiz.Title)
    fmt.Printf("Deskripsi Kuis: %s\n", quiz.Description)

    for _, userQuiz := range userQuizes {
      if userQuiz.QuizID == quiz.ID {
        isFound = true
        fmt.Printf("Nilai: %d\n", userQuiz.Score)
        break
      }
    }

    if !isFound {
      fmt.Println("Nilai: 0")
    }

    fmt.Println()
  }
}