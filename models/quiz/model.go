package quiz

import (
  "modol-app/enums"
  "modol-app/helpers"
)

/**
  * Quiz struct
  */
type Quiz struct {
  ID          int
  Title       string
  Description string
}

// Quizzes collection
var Quizzes []Quiz

/**
  * Initializes the Quizzes collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("quizzes.json") {
    content, err := helpers.ReadFile("quizzes.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &Quizzes)
  } else {
    Quizzes = []Quiz{
      {1, "Quiz 1", "Wawasan Dasar."},
      {2, "Quiz 2", "Wawasan Musik."},
      {3, "Quiz 3", "Wawasan Lanjutan."},
    }

    content, err := helpers.SaveToJSON(Quizzes)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("quizzes.json", content)

    if err != nil {
      panic(err)
    }
  }
}

/**
  * Creates a new ID using insertion sort
  *
  * @return int
  */
func CreateNewID() int {
  LENGTH := len(Quizzes)

  for i := 1; i < LENGTH; i++ {
    key := Quizzes[i]
    j := i - 1

    for j >= 0 && Quizzes[j].ID > key.ID {
      Quizzes[j + 1] = Quizzes[j]
      j = j - 1
    }

    Quizzes[j + 1] = key
  }

  return Quizzes[LENGTH - 1].ID + 1
}

/**
  * Finds a quiz by its ID
  *
  * @param int id
  * @return Quiz
  */
func FindQuizByID(id int) Quiz {
  for _, quiz := range Quizzes {
    if quiz.ID == id {
      return quiz
    }
  }

  return Quiz{}
}

/**
  * Lists all quizzes with selection sort
  *
  * @param enums.SORT sortType
  * @return []Quiz
  */
func ListQuizzes(sortType enums.SORT) []Quiz {
  LENGTH := len(Quizzes)

  for i := 0; i < LENGTH - 1; i++ {
    min := i

    for j := i + 1; j < LENGTH; j++ {
      if sortType == enums.ASC {
        if Quizzes[j].ID < Quizzes[min].ID {
          min = j
        }
      } else {
        if Quizzes[j].ID > Quizzes[min].ID {
          min = j
        }
      }
    }

    Quizzes[i], Quizzes[min] = Quizzes[min], Quizzes[i]
  }

  return Quizzes
}

/**
  * Inserts a quiz
  *
  * @param Quiz quiz
  * @return Quiz
  */
func InsertQuiz(quiz Quiz) Quiz {
  quiz.ID = CreateNewID()
  Quizzes = append(Quizzes, quiz)

  content, err := helpers.SaveToJSON(Quizzes)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("quizzes.json", content)

  if err != nil {
    panic(err)
  }

  return quiz
}

/**
  * Updates a quiz
  *
  * @param int ID
  * @param Quiz quiz
  * @return Quiz
  */
func UpdateQuizByID(ID int, quiz Quiz) Quiz {
  for i, q := range Quizzes {
    if q.ID == ID {
      Quizzes[i] = quiz
    }
  }

  content, err := helpers.SaveToJSON(Quizzes)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("quizzes.json", content)

  if err != nil {
    panic(err)
  }

  return quiz
}

/**
  * Deletes a quiz by its ID
  *
  * @param int quizID
  * @return void
  */
func DeleteQuizByID(quizID int) {
  var newQuizzes []Quiz

  for i := 0; i < len(Quizzes); i++ {
    if Quizzes[i].ID != quizID {
      newQuizzes = append(newQuizzes, Quizzes[i])
    }
  }

  Quizzes = newQuizzes

  content, err := helpers.SaveToJSON(Quizzes)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("quizzes.json", content)

  if err != nil {
    panic(err)
  }
}
