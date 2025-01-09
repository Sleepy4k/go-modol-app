package controllers

import (
	"fmt"
	"modol-app/enums"
	"modol-app/handlers/auth"
	"modol-app/helpers"
	"modol-app/models/assignment"
	"modol-app/models/class_has_user"
	"modol-app/models/discuss"
	"modol-app/models/discuss_reply"
	"modol-app/models/subject"
	"modol-app/models/quiz"
	"modol-app/models/quiz_question"
	"modol-app/models/user_has_quiz"
	"modol-app/models/user_has_assignment"
)

/**
 * ClassController struct
 */
type ClassController struct {}

/**
  * Displays the class menu
  *
  * @return void
  */
func (cc ClassController) Index() {
  for isRunning := false; !isRunning; {
    var choice int

    for choice < 1 || choice > 5 {
      helpers.ClearScreen()
      helpers.DisplayClassMenu()

      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    isUserHasClass := class_has_user.GetAllByUserID(auth.CurrentUser.ID)

    if choice > 0 && choice < 5 && len(isUserHasClass) == 0 {
      fmt.Println("Anda belum terdaftar di kelas manapun.")
      choice = 0
      helpers.ConfirmationScreen()
      continue
    }

    switch choice {
    case 1:
      cc.Materi()
    case 2:
      if auth.IsSiswa(auth.CurrentUser) {
        cc.Tugas()
      } else {
        fmt.Println("Anda tidak memiliki izin untuk mengakses menu ini.")
      }
    case 3:
      if auth.IsSiswa(auth.CurrentUser) {
        cc.Kuis()
      } else {
        fmt.Println("Anda tidak memiliki izin untuk mengakses menu ini.")
      }
    case 4:
      cc.Forum()
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
  * Displays the course materials
  *
  * @return void
  */
func (cc ClassController) Materi() {
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

  choice = 0
  helpers.ClearScreen()

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    fmt.Println("=== List Materi ===")

    subjects := subject.ListSubjects(sortType)
    for _, subject := range subjects {
      fmt.Printf("%d. %s\n", subject.ID, subject.Name)
    }

    subjects_count := len(subjects)

    if subjects_count == 0 {
      fmt.Println("Tidak ada materi yang tersedia.")
      return
    }

    fmt.Println("===================")
    fmt.Println("1. Lihat Materi")
    fmt.Println("2. Kembali")

    for choice < 1 || choice > 3 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    if subjects_count == 0 && choice == 1 {
      fmt.Println("Tidak ada materi yang tersedia.")
      return
    }

    switch choice {
    case 1:
      var ID int

      fmt.Print("ID Materi: ")
      fmt.Scanln(&ID)

      subject := subject.FindSubjectByID(ID)

      helpers.ClearScreen()

      if subject.ID == 0 {
        fmt.Println("Materi tidak ditemukan.")
      } else {
        fmt.Printf("Materi: %s\n", subject.Name)
        fmt.Printf("Deskripsi: %s\n", subject.Description)
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
  * Displays the course assignments
  *
  * @return void
  */
func (cc ClassController) Tugas() {
  fmt.Println("=== Tugas ===")

  var choice int
  helpers.ClearScreen()

  userClass := class_has_user.GetAllByUserID(auth.CurrentUser.ID)

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    fmt.Println("=== List Tugas ===")

    assignments := assignment.GetAllByClassID(userClass[0].ClassID, enums.ASC)
    for _, assignment := range assignments {
      fmt.Printf("%d. %s\n", assignment.ID, assignment.Title)
    }

    assignments_count := len(assignments)

    if assignments_count == 0 {
      fmt.Println("Tidak ada tugas yang tersedia.")
      return
    }

    fmt.Println("===================")
    fmt.Println("1. Lihat Tugas")
    fmt.Println("2. Kumpulkan Tugas")
    fmt.Println("3. Kembali")

    for choice < 1 || choice > 3 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    if assignments_count == 0 && (choice == 1 || choice == 2) {
      fmt.Println("Tidak ada tugas yang tersedia.")
      return
    }

    switch choice {
    case 1:
      var ID int

      fmt.Print("ID Tugas: ")
      fmt.Scanln(&ID)

      assignment := assignment.FindAssignmentByID(ID)

      helpers.ClearScreen()

      if assignment.ID == 0 {
        fmt.Println("Tugas tidak ditemukan.")
      } else {
        fmt.Printf("Tugas: %s\n", assignment.Title)
        fmt.Printf("Deskripsi: %s\n", assignment.Description)
      }
    case 2:
      var ID int
      var submission string

      fmt.Print("ID Tugas: ")
      fmt.Scanln(&ID)

      isAlreadySubmitted := user_has_assignment.GetByUserIDAndAssignmentID(auth.CurrentUser.ID, ID)

      if isAlreadySubmitted.ID != 0 {
        fmt.Println("Anda sudah mengumpulkan tugas ini.")
        return
      }

      assignment := assignment.FindAssignmentByID(ID)

      helpers.ClearScreen()

      if assignment.ID == 0 {
        fmt.Println("Tugas tidak ditemukan.")
      } else {
        fmt.Printf("Tugas: %s\n", assignment.Title)
        fmt.Print("Kumpulkan Tugas: ")
        submission = helpers.GetInlineInput()

        data := user_has_assignment.UserHasAssignment{
          UserID: auth.CurrentUser.ID,
          AssignmentID: assignment.ID,
          Submission: submission,
        }

        user_has_assignment.InsertUserHasAssignment(data)
      }
    case 3:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 3 {
      choice = 0
      helpers.ConfirmationScreen()
    }
  }
}

/**
  * Displays the course quizzes
  *
  * @return void
  */
func (cc ClassController) Kuis() {
  fmt.Println("=== Kuis ===")
  
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

  choice = 0
  helpers.ClearScreen()

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    fmt.Println("=== List Kuis ===")

    quizzes := quiz.ListQuizzes(sortType)
    for _, quiz := range quizzes {
      fmt.Printf("%d. %s\n", quiz.ID, quiz.Title)
    }

    quizzes_count := len(quizzes)

    if quizzes_count == 0 {
      fmt.Println("Tidak ada kuis yang tersedia.")
      return
    }

    fmt.Println("===================")
    fmt.Println("1. Lihat Kuis")
    fmt.Println("2. Kerjakan Kuis")
    fmt.Println("3. Kembali")

    for choice < 1 || choice > 3 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    if quizzes_count == 0 && choice == 1 {
      fmt.Println("Tidak ada kuis yang tersedia.")
      return
    }

    switch choice {
    case 1:
      var ID int

      fmt.Print("ID Kuis: ")
      fmt.Scanln(&ID)

      quiz := quiz.FindQuizByID(ID)

      helpers.ClearScreen()

      if quiz.ID == 0 {
        fmt.Println("Kuis tidak ditemukan.")
      } else {
        quizQuestions := quiz_question.ListAllByQuizID(ID)

        fmt.Printf("Kuis: %s\n", quiz.Title)
        fmt.Printf("Deskripsi: %s\n", quiz.Description)
        fmt.Printf("Jumlah Soal: %d\n", len(quizQuestions))
      }
    case 2:
      var ID, correct int

      fmt.Print("ID Kuis: ")
      fmt.Scanln(&ID)

      isAlreadyDone := user_has_quiz.GetByUserIDAndQuizID(auth.CurrentUser.ID, ID)

      if isAlreadyDone.ID != 0 {
        fmt.Println("Anda sudah mengerjakan kuis ini.")
        return
      }

      quiz := quiz.FindQuizByID(ID)
      helpers.ClearScreen()

      if quiz.ID == 0 {
        fmt.Println("Kuis tidak ditemukan.")
      } else {
        var answer string
        quizQuestions := quiz_question.ListAllByQuizID(ID)

        for i, quizQuestion := range quizQuestions {
          fmt.Printf("=== %s ===\n", quiz.Title)
          fmt.Printf("Soal %d: %s\n", i + 1, quizQuestion.Question)
          fmt.Printf("A. %s\n", quizQuestion.A)
          fmt.Printf("B. %s\n", quizQuestion.B)
          fmt.Printf("C. %s\n", quizQuestion.C)
          fmt.Printf("D. %s\n", quizQuestion.D)
          fmt.Print("Jawaban: ")
          fmt.Scanln(&answer)

          answer = helpers.ToUpper(answer)
          correctAnswer := helpers.ToUpper(quizQuestion.Answer)

          if answer == correctAnswer {
            correct++
          }

          helpers.ClearScreen()
        }

        score := (correct / len(quizQuestions)) * 100

        fmt.Printf("Soal yang benar: %d\n", correct)
        fmt.Printf("Soal yang salah: %d\n", len(quizQuestions) - correct)
        fmt.Printf("Skor: %d\n", score)

        data := user_has_quiz.UserHasQuiz{
          UserID: auth.CurrentUser.ID,
          QuizID: quiz.ID,
          Score: score,
        }

        user_has_quiz.InsertUserHasQuiz(data)
      }
    case 3:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 3 {
      choice = 0
      helpers.ConfirmationScreen()
    }
  }
}

/**
  * Displays the course forums
  *
  * @return void
  */
func (cc ClassController) Forum() {
  var choice int

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    helpers.DisplayForumMenu()

    for choice < 1 || choice > 3 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    switch choice {
    case 1:
      cc.ListForum()
    case 2:
      cc.CreateForum()
    case 3:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 3 {
      choice = 0
      helpers.ConfirmationScreen()
    }
  }
}

/**
  * Displays the list of forums
  *
  * @return void
  */
func (cc ClassController) ListForum() {
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

  choice = 0
  helpers.ClearScreen()

  for isRunning := false; !isRunning; {
    helpers.ClearScreen()
    fmt.Println("=== List Forum ===")

    discusses := discuss.ListDiscusses(sortType)
    for _, discuss := range discusses {
      fmt.Printf("%d. %s\n", discuss.ID, discuss.Content)
    }

    discusses_count := len(discusses)

    if discusses_count == 0 {
      fmt.Println("Tidak ada forum yang tersedia.")
      return
    }

    fmt.Println("===================")
    fmt.Println("1. Lihat Forum")
    fmt.Println("2. Balas Forum")
    fmt.Println("3. Kembali")

    for choice < 1 || choice > 3 {
      fmt.Print("Masukan pilihan anda: ")
      fmt.Scanln(&choice)
    }

    helpers.ClearScreen()

    if discusses_count == 0 && (choice == 1 || choice == 2) {
      fmt.Println("Tidak ada forum yang tersedia.")
      return
    }

    switch choice {
    case 1:
      var ID int

      fmt.Print("ID Forum: ")
      fmt.Scanln(&ID)

      discuss := discuss.FindDiscussByID(ID)

      helpers.ClearScreen()

      if discuss.ID == 0 {
        fmt.Println("Forum tidak ditemukan.")
      } else {
        fmt.Printf("Forum: %s\n", discuss.Content)
        fmt.Println("=== Balasan ===")

        replies := discuss_reply.ListDiscussRepliesByDiscussID(ID, enums.ASC)
        for i, reply := range replies {
          fmt.Printf("%d. %s\n", i + 1, reply.Content)
        }

        if len(replies) == 0 {
          fmt.Println("Belum ada balasan.")
        }
      }
    case 2:
      var ID int
      var content string

      fmt.Print("ID Forum: ")
      fmt.Scanln(&ID)

      discuss := discuss.FindDiscussByID(ID)

      helpers.ClearScreen()

      if discuss.ID == 0 {
        fmt.Println("Forum tidak ditemukan.")
      } else {
        fmt.Printf("Forum: %s\n", discuss.Content)
        fmt.Print("Balasan: ")
        content = helpers.GetInlineInput()

        data := discuss_reply.DiscussReply{
          DiscussID: discuss.ID,
          Content: content,
          UserID: auth.CurrentUser.ID,
        }

        discuss_reply.InsertDiscussReply(data)
      }
    case 3:
      isRunning = true
    default:
      fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }

    if choice != 3 {
      choice = 0
      helpers.ConfirmationScreen()
    }
  }
}

/**
  * Handles the forum creation process
  *
  * @return void
  */
func (cc ClassController) CreateForum() {
  fmt.Println("=== Buat Forum ===")

  var content string

  fmt.Print("Forum: ")
  content = helpers.GetInlineInput()

  data := discuss.Discuss{
    Content: content,
    UserID: auth.CurrentUser.ID,
  }

  discuss.InsertDiscuss(data)

  fmt.Println("Forum berhasil dibuat.")
}
