package user_has_assignment

import (
  "modol-app/helpers"
  "modol-app/models/user"
  "modol-app/models/assignment"
)

/**
  * UserHasAssignment struct
  */
type UserHasAssignment struct {
  ID                    int
  UserID                int
  AssignmentID           int
  Submission            string
}

// UserHasAssignments collection
var UserHasAssignments []UserHasAssignment

/**
  * Initializes the UserHasAssignments collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("user_has_assigments.json") {
    content, err := helpers.ReadFile("user_has_assigments.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &UserHasAssignments)
  } else {
    UserHasAssignments = []UserHasAssignment{
      {1, 3, 1, "https://drive.google.com/file/d/tugas1"},
    }

    content, err := helpers.SaveToJSON(UserHasAssignments)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("user_has_assigments.json", content)

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
  LENGTH := len(UserHasAssignments)

  for i := 1; i < LENGTH; i++ {
    key := UserHasAssignments[i]
    j := i - 1

    for j >= 0 && UserHasAssignments[j].ID > key.ID {
      UserHasAssignments[j + 1] = UserHasAssignments[j]
      j = j - 1
    }

    UserHasAssignments[j + 1] = key
  }

  return UserHasAssignments[LENGTH - 1].ID + 1
}

/**
  * Get the user from the UserHasAssigment
  *
  * @param int id
  * @return user.User
  */
func GetUser(id int) user.User {
  for _, userHasAssigment := range UserHasAssignments {
    if userHasAssigment.ID == id {
      return *user.FindUserByID(userHasAssigment.UserID)
    }
  }

  return user.User{}
}

/**
  * Get the assignment from the UserHasAssigment
  *
  * @param int id
  * @return assignment.Assignment
  */
func GetAssignment(id int) assignment.Assignment {
  for _, userHasAssigment := range UserHasAssignments {
    if userHasAssigment.ID == id {
      return *assignment.FindAssignmentByID(userHasAssigment.AssignmentID)
    }
  }

  return assignment.Assignment{}
}

/**
  * Get all user has assigments
  *
  * @return []UserHasAssignment
  */
func GetAll() []UserHasAssignment {
  return UserHasAssignments
}

/**
  * Get all user has assigments by user ID
  *
  * @param int userID
  * @return []UserHasAssignment
  */
func GetAllByUserID(userID int) []UserHasAssignment {
  var userHasAssignments []UserHasAssignment

  for _, userHasAssignment := range UserHasAssignments {
    if userHasAssignment.UserID == userID {
      userHasAssignments = append(userHasAssignments, userHasAssignment)
    }
  }

  return userHasAssignments
}

/**
  * Get all user has assigments by assignment ID
  *
  * @param int assignmentID
  * @return []UserHasAssignment
  */
func GetByUserIDAndAssignmentID(userID int, assignmentID int) UserHasAssignment {
  for _, userHasAssignment := range UserHasAssignments {
    if userHasAssignment.UserID == userID && userHasAssignment.AssignmentID == assignmentID {
      return userHasAssignment
    }
  }

  return UserHasAssignment{}
}

/**
  * Inserts a user has assigment
  *
  * @param UserHasAssignment userHasAssignment
  * @return void
  */
func InsertUserHasAssignment(userHasAssignment UserHasAssignment) {
  userHasAssignment.ID = CreateNewID()
  UserHasAssignments = append(UserHasAssignments, userHasAssignment)

  content, err := helpers.SaveToJSON(UserHasAssignments)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("user_has_assigments.json", content)

  if err != nil {
    panic(err)
  }
}