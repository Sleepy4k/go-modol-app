package quiz_question

import (
  "modol-app/helpers"
)

/**
  * QuizQuestion struct
  */
type QuizQuestion struct {
  ID          int
  QuizID      int
  Question    string
  A           string
  B           string
  C           string
  D           string
  Answer      string
}

// QuizQuestions collection
var QuizQuestions []QuizQuestion

/**
  * Initializes the QuizQuestions collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("quiz_questions.json") {
    content, err := helpers.ReadFile("quiz_questions.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &QuizQuestions)
  } else {
    QuizQuestions = []QuizQuestion{
      {1, 1, "Apa itu Go?", "Jalan", "Ganesha Operation", "Google", "Golang", "D"},
      {2, 1, "Apa itu Golang?", "Jajanan", "Bocah Petualang", "Obat", "Bahasa Pemrograman", "D"},
      {3, 2, "Apa itu REST API?", "Representational State Transfer", "Resep", "Resolusi", "Resiko", "A"},
      {4, 2, "Apa itu HI-RES?", "Grafik", "Grafis", "Grafis Kualitas Tinggi", "Grafis Kualitas Rendah", "C"},
      {5, 3, "Apa gambaran Docker?", "Kapal", "Kontainer", "Kapal Kontainer", "Kapal Kontainer Besar", "C"},
      {6, 3, "Apa gambaran Kubernetes?", "Kapal", "Kontainer", "Kapal Kontainer", "Kapal Kontainer Besar", "D"},
    }

    content, err := helpers.SaveToJSON(QuizQuestions)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("quiz_questions.json", content)

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
  LENGTH := len(QuizQuestions)

  for i := 1; i < LENGTH; i++ {
    key := QuizQuestions[i]
    j := i - 1

    for j >= 0 && QuizQuestions[j].ID > key.ID {
      QuizQuestions[j + 1] = QuizQuestions[j]
      j = j - 1
    }

    QuizQuestions[j + 1] = key
  }

  return QuizQuestions[LENGTH - 1].ID + 1
}

/**
  * Finds a quiz question by its ID
  *
  * @param int id
  * @return QuizQuestion
  */
func FindQuizQuestionByID(id int) QuizQuestion {
  for _, quizQuestion := range QuizQuestions {
    if quizQuestion.ID == id {
      return quizQuestion
    }
  }

  return QuizQuestion{}
}

/**
  * Get all quiz questions
  *
  * @return []QuizQuestion
  */
func GetAll() []QuizQuestion {
  return QuizQuestions
}

/**
  * Get all quiz questions by quiz ID
  *
  * @param int quizID
  * @return []QuizQuestion
  */
func ListAllByQuizID(quizID int) []QuizQuestion {
  var quizQuestions []QuizQuestion

  for _, quizQuestion := range QuizQuestions {
    if quizQuestion.QuizID == quizID {
      quizQuestions = append(quizQuestions, quizQuestion)
    }
  }

  return quizQuestions
}

/**
  * Inserts a quiz question
  *
  * @param QuizQuestion quizQuestion
  * @return void
  */
func InsertQuizQuestion(quizQuestion QuizQuestion) {
  quizQuestion.ID = CreateNewID()
  QuizQuestions = append(QuizQuestions, quizQuestion)

  content, err := helpers.SaveToJSON(QuizQuestions)

  if err != nil {
    panic(err)
  }

  err = helpers.SaveFile("quiz_questions.json", content)

  if err != nil {
    panic(err)
  }
}

/**
  * Deletes a quiz question by its ID
  *
  * @param int quizQuestionID
  * @return void
  */
func DeleteQuizQuestionByID(quizQuestionID int) {
  var newQuizQuestions []QuizQuestion

  for i := 0; i < len(QuizQuestions); i++ {
    if QuizQuestions[i].ID != quizQuestionID {
      newQuizQuestions = append(newQuizQuestions, QuizQuestions[i])
    }
  }

  QuizQuestions = newQuizQuestions

  content, err := helpers.SaveToJSON(QuizQuestions)

  if err != nil {
    panic(err)
  }

  err = helpers.SaveFile("quiz_questions.json", content)

  if err != nil {
    panic(err)
  }
}