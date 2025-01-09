package user_has_quiz

import "modol-app/helpers"

/**
  * UserHasQuiz struct
  */
type UserHasQuiz struct {
  ID          int
  UserID      int
  QuizID      int
  Score       int
}

// UserHasQuizzes collection
var UserHasQuizzes []UserHasQuiz

/**
  * Initializes the UserHasQuizzes collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("user_has_quizzes.json") {
    content, err := helpers.ReadFile("user_has_quizzes.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &UserHasQuizzes)
  } else {
    UserHasQuizzes = []UserHasQuiz{
      {1, 1, 1, 100},
    }

    content, err := helpers.SaveToJSON(UserHasQuizzes)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("user_has_quizzes.json", content)

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
  LENGTH := len(UserHasQuizzes)

  for i := 1; i < LENGTH; i++ {
    key := UserHasQuizzes[i]
    j := i - 1

    for j >= 0 && UserHasQuizzes[j].ID > key.ID {
      UserHasQuizzes[j + 1] = UserHasQuizzes[j]
      j = j - 1
    }

    UserHasQuizzes[j + 1] = key
  }

  return UserHasQuizzes[LENGTH - 1].ID + 1
}

/**
  * Get all user has quiz by user ID
  * 
  * @param int userID
  * @return []UserHasQuiz
  */
func GetAllByUserID(userID int) []UserHasQuiz {
  var userHasQuizzes []UserHasQuiz

  for _, userHasQuiz := range UserHasQuizzes {
    if userHasQuiz.UserID == userID {
      userHasQuizzes = append(userHasQuizzes, userHasQuiz)
    }
  }

  return userHasQuizzes
}

/**
  * Get all user has quiz by quiz ID
  *
  * @param int quizID
  * @return []UserHasQuiz
  */
func GetAllByQuizID(quizID int) []UserHasQuiz {
  var userHasQuizzes []UserHasQuiz

  for _, userHasQuiz := range UserHasQuizzes {
    if userHasQuiz.QuizID == quizID {
      userHasQuizzes = append(userHasQuizzes, userHasQuiz)
    }
  }

  return userHasQuizzes
}

/**
  * Inserts a user has quiz
  *
  * @param UserHasQuiz userHasQuiz
  * @return void
  */
func InsertUserHasQuiz(userHasQuiz UserHasQuiz) {
  userHasQuiz.ID = CreateNewID()
  UserHasQuizzes = append(UserHasQuizzes, userHasQuiz)

  content, err := helpers.SaveToJSON(UserHasQuizzes)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("user_has_quizzes.json", content)

  if err != nil {
    panic(err)
  }
}

/**
  * Get all user has quiz by user ID and quiz ID
  *
  * @param int userID
  * @param int quizID
  * @return UserHasQuiz
  */
func GetByUserIDAndQuizID(userID, quizID int) UserHasQuiz {
  for _, userHasQuiz := range UserHasQuizzes {
    if userHasQuiz.UserID == userID && userHasQuiz.QuizID == quizID {
      return userHasQuiz
    }
  }

  return UserHasQuiz{}
}
